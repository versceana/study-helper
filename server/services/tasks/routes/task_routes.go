package routes

import (
	"server/services/tasks/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine, controller *controllers.TaskController) {
	if router == nil {
		panic("router is nil")
	}

	if controller == nil {
		panic("controller is nil")
	}

	tasks := router.Group("/tasks")
	{
		tasks.GET("/", controller.GetTasks)
		tasks.POST("/", controller.AddTasks)
		tasks.DELETE("/:id", controller.DeleteTask)
	}
}