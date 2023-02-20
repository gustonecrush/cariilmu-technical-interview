package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gustonecrush/cariilmu-technical-interview/app/models"
	"github.com/gustonecrush/cariilmu-technical-interview/helper"
)

func (server *Server) Index(w http.ResponseWriter, r *http.Request)  {
	var categories []models.CourseCategory

	if err := server.DB.Find(&categories).Error; err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseJSON(w, http.StatusOK, categories)
}

func (server *Server) Store(w http.ResponseWriter, r *http.Request) {
	// mengambil input json
	var userInput models.CourseCategory
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

func (server *Server) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	var category models.CourseCategory
	
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		helper.ResponseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	if server.DB.Where("id = ?", id).Updates(&category).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, "could not update category")
		return
	}

	category.ID = id

	helper.ResponseJSON(w, http.StatusOK, category)
}

func (server *Server) Destroy(w http.ResponseWriter, r *http.Request) {
	// megambil input 
	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		helper.ResponseJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	var category models.CourseCategory
	if server.DB.Delete(&category, input["id"]).RowsAffected == 0 {
		helper.ResponseJSON(w, http.StatusBadRequest, "could not delete category")
		return
	}

	response := map[string]string{"message": "category successfully deleted"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

