package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	http.Server
}

func NewServer(port uint16) (s *Server) {
	s = new(Server)
	s.Addr = fmt.Sprintf(":%d", port)
	s.Handler = setRouter()
	return
}

func setRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/login", login)
	r.Get("/verify", verify)

	r.Post("/registration", signIn)

	return r
}

func (s *Server) Start() {
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server running error: %s\n", err.Error())
	}
}
