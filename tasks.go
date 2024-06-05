package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type TasksService struct {
	store Store
}

func NewTasksService(s Store) *TasksService {
	return &TasksService{store: s}
}
func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.handleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.handleGetTask).Methods("GET")
}
func (s *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {

}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {

}
