package controllers

import (
	"github.com/gorilla/mux"
)

func (server *Server) InitializedRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/api/login", server.Login).Methods("POST")
	server.Router.HandleFunc("/api/register", server.Register).Methods("POST")
	server.Router.HandleFunc("/api/logout", Logout).Methods("GET")
}