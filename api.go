package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	addr  string
	store Store
}

func NewApiServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Registering Services
	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(router)

	log.Println("Starting API Server at ", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
