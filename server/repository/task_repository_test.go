package repository

import (
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetAllTasks(t *testing.T) {
	t.Run("when repository DB is nil", func(t *testing.T) {
		repo := &TaskRepository{}
		tasks, err := repo.GetAllTasks()
		assert.Nil(t, tasks)
		assert.Error(t, err)
		assert.Equal(t, "repository DB is nil", err.Error())
	})

	t.Run("when repository DB is not nil, but Find returns an error", func(t *testing.T) {
		dsn := "host=localhost user=correctuser password=correctpassword dbname=correctdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		assert.NoError(t, err)

		repo := &TaskRepository{DB: db}
		tasks, err := repo.GetAllTasks()
		assert.Nil(t, tasks)
		assert.Error(t, err)
		assert.Equal(t, "pq: relation \"tasks\" does not exist", err.Error())
	})

	t.Run("when repository DB is not nil and Find returns no error", func(t *testing.T) {
		dsn := "host=localhost user=correctuser password=correctpassword dbname=correctdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		assert.NoError(t, err)

		// Create table if it doesn't exist
		db.AutoMigrate(&models.Task{})

		// Insert test data
		t1 := models.Task{Description: "task1"}
		t2 := models.Task{Description: "task2"}
		db.Create(&t1)
		db.Create(&t2)

		repo := &TaskRepository{DB: db}
		tasks, err := repo.GetAllTasks()
		assert.NoError(t, err)
		assert.Len(t, tasks, 2)
		assert.NotNil(t, tasks[0])
		assert.NotNil(t, tasks[1])
	})
}
