package controllers

import "github.com/alexymumo/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/note", middlewares.SetMiddlewareJSON(s.CreateNote)).Methods("POST")
	s.Router.HandleFunc("/notes", middlewares.SetMiddlewareJSON(s.FindNotes)).Methods("GET")
	s.Router.HandleFunc("/notes/{id}", middlewares.SetMiddlewareJSON(s.DeleteNote)).Methods("DELETE")
	s.Router.HandleFunc("/note/{id}", middlewares.SetMiddlewareJSON(s.UpdateNote)).Methods("PUT")
}
