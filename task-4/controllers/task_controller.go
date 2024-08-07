package controllers

import (
	"backend-learning-track/task-4/router"

	"github.com/gin-gonic/gin"
)

func Runner(r *gin.Engine) {
	router.SetUpRouter(r)
	r.Run("localhost:8080")

}


