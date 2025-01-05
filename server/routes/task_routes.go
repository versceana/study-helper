package routes

import (
	"server/handlers"

	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(r *mux.Router, handler *handlers.TaskHandler) {
	r.HandleFunc("/tasks", handler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
}
