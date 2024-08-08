package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title,omitempty" json:"title"`
	Description string             `bson:"description,omitempty" json:"description"`
	DueDate     string             `bson:"due_date,omitempty" json:"due_date"`
	Status      bool               `bson:"status,omitempty" json:"status"`
}
