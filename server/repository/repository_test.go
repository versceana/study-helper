package repository

import (
	"os"
	"server/models"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB *gorm.DB

// Настройка тестовой базы данных
func TestMain(m *testing.M) {
	dsn := "host=localhost port=5432 user=postgres password=password dbname=my_database sslmode=disable"
	var err error
	testDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}

	// Миграция моделей
	err = testDB.AutoMigrate(&models.Event{}, &models.Task{})
	if err != nil {
		panic("Failed to migrate test database")
	}

	code := m.Run()

	// Удаление таблиц после тестов
	testDB.Exec("DROP TABLE events;")
	testDB.Exec("DROP TABLE tasks;")

	os.Exit(code)
}

// Тесты для EventRepository
func TestEventRepository(t *testing.T) {
	repo := &EventRepository{DB: testDB}

	// Тест на сохранение события
	event := models.Event{
		GoogleID:    "test-google-id",
		Title:       "Test Event",
		Description: "This is a test event",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(2 * time.Hour),
	}
	err := repo.SaveEvent(&event)
	if err != nil {
		t.Fatalf("Failed to save event: %v", err)
	}

	// Тест на получение всех событий
	events, err := repo.GetAllEvents()
	if err != nil {
		t.Fatalf("Failed to get all events: %v", err)
	}
	if len(events) != 1 {
		t.Errorf("Expected 1 event, got %d", len(events))
	}

	// Тест на удаление события
	err = repo.DeleteEvent(event.ID)
	if err != nil {
		t.Fatalf("Failed to delete event: %v", err)
	}

	// Проверка отсутствия событий
	events, err = repo.GetAllEvents()
	if err != nil {
		t.Fatalf("Failed to get all events after deletion: %v", err)
	}
	if len(events) != 0 {
		t.Errorf("Expected 0 events, got %d", len(events))
	}
}

// Тесты для TaskRepository
func TestTaskRepository(t *testing.T) {
	repo := &TaskRepository{DB: testDB}

	// Тест на создание задачи
	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Completed:   false,
	}
	err := repo.CreateTask(&task)
	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	// Тест на получение всех задач
	tasks, err := repo.GetAllTasks()
	if err != nil {
		t.Fatalf("Failed to get all tasks: %v", err)
	}
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}

	// Тест на удаление задачи
	err = repo.DeleteTask(task.ID)
	if err != nil {
		t.Fatalf("Failed to delete task: %v", err)
	}

	// Проверка отсутствия задач
	tasks, err = repo.GetAllTasks()
	if err != nil {
		t.Fatalf("Failed to get all tasks after deletion: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}
}
