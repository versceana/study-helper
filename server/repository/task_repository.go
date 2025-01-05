package repository

import (
	"fmt"
	"gorm.io/gorm"
	"server/models"
)

type TaskRepository struct {
	DB *gorm.DB
}

func (repo *TaskRepository) GetAllTasks() ([]models.Task, error) {
	if repo.DB == nil {
		return nil, fmt.Errorf("repository DB is nil")
	}

	var tasks []models.Task
	err := repo.DB.Find(&tasks).Error
	return tasks, err
}

func (repo *TaskRepository) CreateTask(task *models.Task) error {
	if repo.DB == nil {
		return fmt.Errorf("repository DB is nil")
	}
	if task == nil {
		return fmt.Errorf("task is nil")
	}
	return repo.DB.Create(task).Error
}

func (repo *TaskRepository) DeleteTask(id uint) error {
	if repo.DB == nil {
		return fmt.Errorf("repository DB is nil")
	}
	result := repo.DB.Delete(&models.Task{}, id)
	if result.Error != nil {
		return fmt.Errorf("error deleting task: %w", result.Error)
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("task with id %d does not exist", id)
	}
	return nil
}
