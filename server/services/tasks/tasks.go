package tasks

import (
    "server/services/tasks/controllers"
    "server/services/tasks/repository"
    "server/services/tasks/routes"

    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
)

func RegisterTasksModule(router *gin.Engine, db *gorm.DB) {
	taskRepo := &repository.TaskRepository{DB: db}
	taskController := &controllers.TaskController{Repo: taskRepo}
	routes.RegisterTaskRoutes(router, taskController)
}
