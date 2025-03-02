package server

import (
	"log"
	"net/http"

	"github.com/AlexDillz/calc2sprint/internal/config"
	"github.com/gorilla/mux"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/calculate", CalculateHandler).Methods("POST")
	router.HandleFunc("/api/v1/expressions", GetAllExpressionsHandler).Methods("GET")
	router.HandleFunc("/api/v1/expressions/{id:[0-9]+}", GetExpressionHandler).Methods("GET")

	router.HandleFunc("/internal/task", TaskHandler).Methods("GET", "POST")

	addr := s.cfg.ServerAddr + ":" + s.cfg.Port
	log.Printf("Server is running on %s", addr)
	return http.ListenAndServe(addr, router)
}
