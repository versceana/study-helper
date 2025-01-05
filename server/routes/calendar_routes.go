package routes

import (
	"server/handlers"

	"github.com/gorilla/mux"
)

func RegisterCalendarRoutes(r *mux.Router, handler *handlers.CalendarHandler) {
	r.HandleFunc("/auth/google", handler.GoogleLogin).Methods("GET")
	r.HandleFunc("/auth/callback", handler.GoogleCallback).Methods("GET")
}
