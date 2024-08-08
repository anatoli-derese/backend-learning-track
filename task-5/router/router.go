package router

import (
	"backend-learning-track/task-5/controllers"

	"github.com/gin-gonic/gin"
)

type Router interface {
	SetUpRouter(router *gin.Engine)
}
type RouterImp struct {
	controller controllers.TaskControllerImp
}

func NewRouter(controller controllers.TaskControllerImp) *RouterImp {

	return &RouterImp{controller: controller}
}

func (c *RouterImp) SetUpRouter(router *gin.Engine) {
	router.GET("/tasks", c.controller.GetAllTasks)
	router.GET("/tasks/:id", c.controller.GetSpecificTask)
	router.POST("/tasks", c.controller.AddNewTask)
	router.PUT("/tasks/:id", c.controller.UpdateTask)
	router.DELETE("/tasks/:id", c.controller.DeleteTask)
	router.Run("localhost:8080")

}
