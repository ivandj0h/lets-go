package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	api := NewApiServer(":3000", nil)
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Registering Services
	log.Println("Starting API Server at ", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
