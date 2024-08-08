package data

import (
	"backend-learning-track/task-4/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var taskID = 1

var tasks []models.Task

func GetAllTask(c *gin.Context) {
	if len(tasks) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "There are no tasks at the moment! Use the Post Method to add a new task."})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetSpecificTask(intId int, c *gin.Context) {
	for _, task := range tasks {
		if task.ID == intId {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	resp := strings.Join([]string{"Task with ID ", strconv.Itoa(intId), " not found."}, "")
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": resp})
}

func AddNewTask(newTask models.Task, c *gin.Context) {
	newTask.ID = taskID
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func DeleteTask(intId int, c *gin.Context) {

	for i, task := range tasks {
		if task.ID == intId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Deleted"})
			return
		}
	}
	resp := strings.Join([]string{"Task with ID ", strconv.Itoa(intId), " not found."}, "")
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": resp})

}

func UpdateTask(intId int, updatedTask models.Task, c *gin.Context) {
	for i, task := range tasks {
		if task.ID == intId {
			tasks[i] = updatedTask
			tasks[i].ID = intId
			c.IndentedJSON(http.StatusOK, tasks[i])
			return
		}
	}
	updatedTask.ID = taskID
	c.IndentedJSON(http.StatusCreated, updatedTask)
	taskID++

}
