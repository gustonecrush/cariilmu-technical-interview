package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (server *Server) DestroyCourse(w http.ResponseWriter, r *http.Request) {
	// mangambil input
	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	var course models.Course
	if server.DB.Delete(&course, input["id"]).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, "could not delete course")
		return
	}

	response := map[string]string{"message": "course successfully deleted"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func (server *Server) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	var course models.Course

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&course); err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if server.DB.Where("id = ?", id).Updates(&course).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, "could not updated course")
		return
	}

	course.ID = id

	helper.ResponseJSON(w, http.StatusOK, course)
}