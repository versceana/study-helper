package controllers

import (
	"net/http"
	"strconv"
	"server/services/tasks/models"
	"server/services/tasks/repository"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Repo *repository.TaskRepository
}

func (ctrl *TaskController) GetTasks(c *gin.Context) {
	tasks, err := ctrl.Repo.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (ctrl *TaskController) AddTasks(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ctrl.Repo == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "repository not initialized"})
		return
	}

	if err := ctrl.Repo.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (ctrl *TaskController) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Repo.DeleteTask(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
