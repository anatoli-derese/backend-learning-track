package controllers

import (
	"backend-learning-track/task-4/data"

	"github.com/gin-gonic/gin"
)

func Runner(router *gin.Engine) {
	router.GET("/tasks", data.GetAllTask)
	router.GET("/tasks/:id", data.GetSpecificTask)
	router.POST("/tasks", data.AddNewTask)
	router.DELETE("/tasks/:id", data.DeleteTask)
	router.PUT("/tasks/:id", data.UpdateTask)
	router.Run("localhost:8080")

}
