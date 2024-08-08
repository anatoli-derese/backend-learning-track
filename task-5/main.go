package main

import (
	"backend-learning-track/task-4/router"
	"context"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	router := gin.Default()
	
	Runner(router)
}
func DataBaseConnection() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("test").Collection("trainers")


}

func Runner(r *gin.Engine) {
	router.SetUpRouter(r)
	r.Run("localhost:8080")

}
