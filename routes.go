package main

import (
	"delayed-queue/task"

	"github.com/gorilla/mux"
)

func newRouter(service task.Service) *mux.Router {

	handler := task.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/create", handler.Create).Methods("POST")
	r.HandleFunc("/tasks", handler.Tasks).Methods("GET")
	r.HandleFunc("/task/{id}", handler.Task).Methods("GET")
	return r
}
