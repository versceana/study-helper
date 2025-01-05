package services

import (
	"server/models"
	"server/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.Repo.GetAllTasks()
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.Repo.CreateTask(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.Repo.DeleteTask(id)
}
