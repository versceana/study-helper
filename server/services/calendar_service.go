package services

import (
	"context"
	"fmt"
	"server/models"
	"server/repository"
	"server/auth"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

type CalendarService struct {
	Repo *repository.EventRepository
}

func (s *CalendarService) SyncGoogleCalendar(token *oauth2.Token) error {
	// Инициализация клиента Google Calendar API
	client := auth.GoogleOAuthConfig.Client(context.Background(), token)
	service, err := calendar.New(client)
	if err != nil {
		return err
	}

	// Получение списка событий
	events, err := service.Events.List("primary").Do()
	if err != nil {
		return err
	}

	// Сохранение событий в базе данных
	for _, e := range events.Items {
		// Парсинг строковых значений в time.Time
		startTime, err := time.Parse(time.RFC3339, e.Start.DateTime)
		if err != nil {
			return fmt.Errorf("failed to parse start time: %w", err)
		}
		endTime, err := time.Parse(time.RFC3339, e.End.DateTime)
		if err != nil {
			return fmt.Errorf("failed to parse end time: %w", err)
		}

		event := &models.Event{
			GoogleID:    e.Id,
			Title:       e.Summary,
			Description: e.Description,
			StartTime:   startTime,
			EndTime:     endTime,
		}

		// Сохранение события в базе данных
		if err := s.Repo.SaveEvent(event); err != nil {
			return err
		}
	}
	return nil
}
