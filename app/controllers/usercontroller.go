package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gustonecrush/cariilmu-technical-interview/app/models"
	"github.com/gustonecrush/cariilmu-technical-interview/helper"
)

func (server *Server) IndexUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	if err := server.DB.Find(&users).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseJSON(w, http.StatusOK, users)
}

func (server *Server) StoreUser(w http.ResponseWriter, r *http.Request) {
	// mengambil input json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	} 
	defer r.Body.Close()

	// insert to database
	if err := server.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if server.DB.Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, "could not update user")
		return
	}

	user.Id = id

	helper.ResponseJSON(w, http.StatusOK, user)
}

func (server *Server) DestroyUser(w http.ResponseWriter, r *http.Request) {
	// mengambil input
	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	var user models.User
	if server.DB.Delete(&user, input["id"]).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, "could not update user")
		return
	}

	response := map[string]string{"message": "user successfully deleted"}
	helper.ResponseJSON(w, http.StatusOK, response)
}