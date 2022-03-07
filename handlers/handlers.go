package handlers

import (
	"net/http"
)

//IEventHandler is implemented by handlers
type IEventHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	UpdateDetails(w http.ResponseWriter, r *http.Request)
	Cancel(w http.ResponseWriter, r *http.Request)
	Reschedule(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

func NewEventHandler() IEventHandler {
	return &handler{}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	panic("Implement Me")
}

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	panic("Implement Me")
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	panic("Implement Me")
}

func (h *handler) UpdateDetails(w http.ResponseWriter, r *http.Request) {
	panic("Implement Me")
}

func (h *handler) Cancel(w http.ResponseWriter, r *http.Request) {
	panic("Implement Me")
}

func (h *handler) Reschedule(w http.ResponseWriter, r *http.Request) {
	panic("Implement Me")
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("Implement Me")
}
