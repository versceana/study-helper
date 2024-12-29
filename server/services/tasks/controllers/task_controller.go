package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"server/services/tasks/models"
	"server/services/tasks/repository"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Repo *repository.TaskRepository
}

func NewTaskController(repo *repository.TaskRepository) *TaskController {
	return &TaskController{Repo: repo}
}

func (ctrl *TaskController) GetTasks(c *gin.Context) {
	tasks, err := ctrl.Repo.GetAllTasks()
	if err != nil {
		ctrl.handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (ctrl *TaskController) AddTasks(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		ctrl.handleError(c, err)
		return
	}
	if err := ctrl.validateTask(&task); err != nil {
		ctrl.handleError(c, err)
		return
	}
	if err := ctrl.Repo.CreateTask(&task); err != nil {
		ctrl.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (ctrl *TaskController) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctrl.handleError(c, err)
		return
	}
	if err := ctrl.Repo.DeleteTask(uint(id)); err != nil {
		ctrl.handleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (ctrl *TaskController) handleError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func (ctrl *TaskController) validateTask(task *models.Task) error {
	if task.Text == "" {
		return fmt.Errorf("task text is empty")
	}
	return nil
}
