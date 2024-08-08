package controllers

import (
	"backend-learning-track/task-5/data"
	"backend-learning-track/task-5/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController interface {
	AddNewTask(c *gin.Context)
	GetSpecificTask(c *gin.Context)
	GetAllTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	UpdateTask(c *gin.Context)
}

type TaskControllerImp struct {
	service data.TaskService
}

func NewTaskController(service data.TaskService) *TaskControllerImp {
	return &TaskControllerImp{service: service}
}

func (controller *TaskControllerImp) AddNewTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Task"})
		return
	}
	task, err := controller.service.AddNewTask(newTask)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusCreated, task)
}

func (controller *TaskControllerImp) GetSpecificTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	task, err := controller.service.GetSpecificTask(intId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (controller *TaskControllerImp) GetAllTasks(c *gin.Context) {
	tasks, err := controller.service.GetAllTasks()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No tasks found, please add a task"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func (controller *TaskControllerImp) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	error := controller.service.DeleteTask(intId)
	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (controller *TaskControllerImp) UpdateTask(c *gin.Context) {

	id := c.Param("id")
	intId, err := primitive.ObjectIDFromHex(id)
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

	error := controller.service.UpdateTask(updatedTask)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": error})
		return
	}
	c.IndentedJSON(http.StatusAccepted, updatedTask)

}
