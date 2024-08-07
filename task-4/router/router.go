package router

import (
	"backend-learning-track/task-4/data"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(router *gin.Engine) {
	router.GET("/tasks", data.GetAllTask)
	router.GET("/tasks/:id", data.GetSpecificTask)
	router.POST("/tasks", data.AddNewTask)
	router.PUT("/tasks/:id", data.UpdateTask)
	router.DELETE("/tasks/:id", data.DeleteTask)

}
