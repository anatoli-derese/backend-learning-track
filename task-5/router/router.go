package router

import (
	"backend-learning-track/task-4/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(router *gin.Engine) {
	router.GET("/tasks", controllers.GetAllTask)
	router.GET("/tasks/:id", controllers.GetSpecificTask)
	router.POST("/tasks", controllers.AddNewTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

}
