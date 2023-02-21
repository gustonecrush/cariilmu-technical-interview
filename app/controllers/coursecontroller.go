package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gustonecrush/cariilmu-technical-interview/app/models"
	"github.com/gustonecrush/cariilmu-technical-interview/helper"
)

func (server *Server) IndexCourse(w http.ResponseWriter, r *http.Request) {
	var courses []models.Course

	if err := server.DB.Find(&courses).Error; err != nil {
		var courses []models.Course

		if err := server.DB.Find(&courses).Error; err != nil {
			helper.ResponseJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		helper.ResponseJSON(w, http.StatusOK, courses)
	}
}

func (server *Server) StoreCourse(w http.ResponseWriter, r *http.Request) {
	// megambil input json
	var userInput models.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {

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

func (server *Server) UpdateCourse(w http.ResponseWriter, r *http.Request) {

}

func (server *Server) DeleteCourse(w http.ResponseWriter, r *http.Request) {

}