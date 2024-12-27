package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Text string `gorm:"type:text"`
	Completed bool `gorm:"default:false"`
}
