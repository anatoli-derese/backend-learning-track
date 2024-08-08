package controllers

import (
	"backend-learning-track/task-4/data"
	"backend-learning-track/task-4/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var taskID int = 1

func AddNewTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		fmt.Println(newTask)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task"})
		return
	}
	newTask.ID = taskID
	data.AddNewTask(newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
	taskID++
}

func GetSpecificTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	task, err := data.GetSpecificTask(intId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func GetAllTask(c *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No tasks found, please add a task"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)

}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	error := data.DeleteTask(intId)
	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
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
	updatedTask.ID = intId
	code, err := data.UpdateTask(updatedTask)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Operation Failed"})
	}
	c.IndentedJSON(code, updatedTask)

}
