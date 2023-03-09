package controller

import (
	"golang-crud/middlewares"
)

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")

	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods("GET")

	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateUser))).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.DeleteUser)).Methods("DELETE")

}
