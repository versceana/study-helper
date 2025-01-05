package models

import "time"

type Event struct {
    ID          uint      `gorm:"primaryKey"`
    GoogleID    string    `gorm:"unique;not null"`
    Title       string    `gorm:"type:varchar(255);not null"`
    Description string    `gorm:"type:text"`
    StartTime   time.Time `gorm:"type:timestamp;not null"`
    EndTime     time.Time `gorm:"type:timestamp;not null"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}