package main

import (
	"backend-learning-track/task-5/controllers"
	"backend-learning-track/task-5/data"
	"backend-learning-track/task-5/router"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db, err := DataBaseConnector()

	if err != nil {
		fmt.Println("Error while connecting to database")
		return
	}

	service := data.NewTaskService(db, context.TODO())
	controller := controllers.NewTaskController(service)
	router := router.NewRouter(*controller)
	router.SetUpRouter(gin.Default())

	fmt.Print(router)

	fmt.Print(db)

}

func DataBaseConnector() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	collection := client.Database("Database").Collection("Tasks")
	fmt.Println("Connected to MongoDB")
	return *&collection, nil
}
