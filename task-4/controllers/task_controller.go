package controllers

import (
	"backend-learning-track/task-4/data"
	"backend-learning-track/task-4/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddNewTask(c *gin.Context) {
	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		fmt.Println(newTask)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task"})
		return
	}
	data.AddNewTask(newTask, c)
}

func GetSpecificTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	data.GetSpecificTask(intId, c)
}

func GetAllTask(c *gin.Context) {
	data.GetAllTask(c)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	data.DeleteTask(intId, c)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task"})
		return
	}
	data.UpdateTask(intId, updatedTask, c)
}
