package controllers

import (
	"github.com/gorilla/mux"
)

func (server *Server) InitializedRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/login", server.Login).Methods("POST")
	server.Router.HandleFunc("/register", server.Register).Methods("POST")
	server.Router.HandleFunc("/logout", Logout).Methods("GET")
}