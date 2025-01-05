package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	DueDate     time.Time `gorm:"type:timestamp"`
	Completed   bool      `gorm:"default:false"`
}

func (t *Task) UpdateDescription(newDescription string) {
	t.Description = newDescription
}
