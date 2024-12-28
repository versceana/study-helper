package repository

import (
	"gorm.io/gorm"
	"server/services/tasks/models"
)

type TaskRepository struct {
	DB *gorm.DB
}

func (repo *TaskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks [] models.Task
	err := repo.DB.Find(&tasks).Error
	return tasks, err
}

func (repo *TaskRepository) CreateTask(task *models.Task) error {
	return repo.DB.Create(task).Error
}

func (repo *TaskRepository) DeleteTask(id uint) error {
	return repo.DB.Delete(&models.Task{}, id).Error
}