package main

import (
	"backend-learning-track/task-4/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	controllers.Runner(router)

}
