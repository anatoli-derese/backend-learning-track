package data

import (
	"backend-learning-track/task-4/models"
	"fmt"
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

func GetSpecificTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	for _, task := range tasks {
		if task.ID == intId {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	resp := strings.Join([]string{"Task with ID ", id, " not found."}, "")
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": resp})

}

func AddNewTask(c *gin.Context) {
	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		fmt.Println(newTask)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task"})
		return
	}
	newTask.ID = taskID

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	for i, task := range tasks {
		if task.ID == intId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Deleted"})
			return
		}
	}
	resp := strings.Join([]string{"Task with ID ", id, " not found."}, "")
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": resp})

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
