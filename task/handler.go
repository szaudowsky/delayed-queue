package task

import "net/http"

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Tasks(w http.ResponseWriter, r *http.Request)
	Task(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
}

func NewHandler(taskService Service) Handler {
	return &handler{service: taskService}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) Tasks(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) Task(w http.ResponseWriter, r *http.Request) {

}
