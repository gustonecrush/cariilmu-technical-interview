package controllers

import (
	"github.com/gorilla/mux"
)

func (server *Server) InitializedRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/api/login", server.Login).Methods("POST")
	server.Router.HandleFunc("/api/register", server.Register).Methods("POST")
	server.Router.HandleFunc("/api/logout", server.Logout).Methods("GET")

	server.Router.HandleFunc("/api/course-categories", server.Index).Methods("GET")
	server.Router.HandleFunc("/api/course-categories", server.Store).Methods("POST")
	server.Router.HandleFunc("/api/course-categories/{id}", server.Update).Methods("PUT")
	server.Router.HandleFunc("/api/course-categories", server.Destroy).Methods("DELETE")
}