package handlers

import (
	"net/http"
	"server/auth"
	"server/services"
)

type CalendarHandler struct {
	Service *services.CalendarService
}

func (h *CalendarHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := auth.GetGoogleLoginURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *CalendarHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token, err := auth.ExchangeCodeForToken(code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	err = h.Service.SyncGoogleCalendar(token)
	if err != nil {
		http.Error(w, "Failed to sync calendar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Calendar synchronized successfully"))
}
