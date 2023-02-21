package controllers

import (
	"github.com/gorilla/mux"
)

func (server *Server) InitializedRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/api/login", server.Login).Methods("POST")
	server.Router.HandleFunc("/api/register", server.Register).Methods("POST")
	server.Router.HandleFunc("/api/logout", server.Logout).Methods("GET")

	server.Router.HandleFunc("/api/users", server.IndexUser).Methods("GET")
	server.Router.HandleFunc("/api/users", server.StoreUser).Methods("POST")
	server.Router.HandleFunc("/api/users/{id}", server.UpdateUser).Methods("PUT")
	server.Router.HandleFunc("/api/users", server.DestroyUser).Methods("DELETE")

	server.Router.HandleFunc("/api/course-categories", server.Index).Methods("GET")
	server.Router.HandleFunc("/api/course-categories", server.Store).Methods("POST")
	server.Router.HandleFunc("/api/course-categories/{id}", server.Update).Methods("PUT")
	server.Router.HandleFunc("/api/course-categories", server.Destroy).Methods("DELETE")

	server.Router.HandleFunc("/api/courses", server.IndexCourse).Methods("GET")
	server.Router.HandleFunc("/api/courses", server.StoreCourse).Methods("POST")
	server.Router.HandleFunc("/api/courses", server.UpdateCourse).Methods("PUT")
	server.Router.HandleFunc("/api/courses", server.DestroyCourse).Methods("DELETE")
}