package main

import (
	"backend-learning-track/task-3/controllers"
	"backend-learning-track/task-3/services"
	// "backend-learning-track/task-3/models"
	// "backend-learning-track/task-3/services"
	// "fmt"
)

func main() {

	myLibrary := services.NewLibrary()
	
	controllers.TakeInputAndDelegate(myLibrary)

}
