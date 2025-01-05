package repository

import (
	"server/models"

	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func (r *EventRepository) SaveEvent(event *models.Event) error {
	return r.DB.Create(event).Error
}

func (r *EventRepository) GetAllEvents() ([]models.Event, error) {
	var events []models.Event
	err := r.DB.Find(&events).Error
	return events, err
}

func (r *EventRepository) DeleteEvent(id uint) error {
	return r.DB.Delete(&models.Event{}, id).Error
}
