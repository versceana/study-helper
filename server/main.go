package main

import (
	"log"
	"net/http"
	"server/auth"
	"server/config"
	"server/database"
	"server/handlers"
	"server/repository"
	"server/routes"
	"server/services"

	"github.com/gorilla/mux"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Подключение к базе данных
	database := database.ConnectDatabase(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Инициализация OAuth2
	auth.InitGoogleOAuth(cfg.GoogleClientID, cfg.GoogleClientSecret, cfg.RedirectURL)

	// Инициализация зависимостей
	eventRepo := &repository.EventRepository{DB: database}
	calendarService := &services.CalendarService{Repo: eventRepo}
	calendarHandler := &handlers.CalendarHandler{Service: calendarService}

	// Настройка маршрутов
	router := mux.NewRouter()
	routes.RegisterCalendarRoutes(router, calendarHandler)

	// Запуск сервера
	log.Printf("Server started on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
