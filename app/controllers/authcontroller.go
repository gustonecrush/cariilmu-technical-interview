package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gustonecrush/cariilmu-technical-interview/app/models"
	"github.com/gustonecrush/cariilmu-technical-interview/config"
	"github.com/gustonecrush/cariilmu-technical-interview/helper"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	// mengambil input json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// ambil data user berdasarkan email
	var user models.User
	if err := server.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		switch  err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Email or password is wrong"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// cek apakah password valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Email or password is wrong"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims  := &config.JWTClaim {
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims {
			Issuer: "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasikan algoritma yang akan digunakan untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token 
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": "Email or password is wrong"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set token yang ke cookie
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Login successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func (server *Server) Register(w http.ResponseWriter, r *http.Request) {
	// mengambil input json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// hash password using bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert to database
	if err := server.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {

}