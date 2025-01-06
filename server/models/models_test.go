package models

import (
	"os"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	dsn := "host=localhost port=5432 user=postgres password=password dbname=my_database sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}

	err = db.AutoMigrate(&Event{}, &Task{})
	if err != nil {
		panic("Failed to migrate test database")
	}

	code := m.Run()

	db.Exec("DROP TABLE events;")
	db.Exec("DROP TABLE tasks;")

	os.Exit(code)
}

func TestEventModel(t *testing.T) {
	event := Event{
		GoogleID:    "test-google-id",
		Title:       "Test Event",
		Description: "This is a test event",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(2 * time.Hour),
	}

	if err := db.Create(&event).Error; err != nil {
		t.Fatalf("Failed to create event: %v", err)
	}

	var fetchedEvent Event
	if err := db.First(&fetchedEvent, "google_id = ?", "test-google-id").Error; err != nil {
		t.Fatalf("Failed to fetch event: %v", err)
	}

	if fetchedEvent.Title != event.Title {
		t.Errorf("Expected title %s, got %s", event.Title, fetchedEvent.Title)
	}

	if err := db.Delete(&fetchedEvent).Error; err != nil {
		t.Fatalf("Failed to delete event: %v", err)
	}
}

func TestTaskModel(t *testing.T) {
	task := Task{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Completed:   false,
	}

	if err := db.Create(&task).Error; err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	newDescription := "Updated task description"
	task.Description = newDescription
	if err := db.Save(&task).Error; err != nil {
		t.Fatalf("Failed to update task: %v", err)
	}

	var fetchedTask Task
	if err := db.First(&fetchedTask, task.ID).Error; err != nil {
		t.Fatalf("Failed to fetch task: %v", err)
	}

	if fetchedTask.Description != newDescription {
		t.Errorf("Expected description %s, got %s", newDescription, fetchedTask.Description)
	}

	if err := db.Delete(&fetchedTask).Error; err != nil {
		t.Fatalf("Failed to delete task: %v", err)
	}
}
