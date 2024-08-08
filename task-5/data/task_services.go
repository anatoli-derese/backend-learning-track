package data

import (
	"backend-learning-track/task-5/models"
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService interface {
	GetAllTasks() ([]models.Task, error)
	GetSpecificTask(intId primitive.ObjectID) (models.Task, error)
	AddNewTask(newTask models.Task) (models.Task, error)
	DeleteTask(intId primitive.ObjectID) error
	UpdateTask(updatedTask models.Task) error
}

type TaskServiceStruct struct {
	db  *mongo.Collection
	ctx context.Context
}

func NewTaskService(db *mongo.Collection, ctx context.Context) TaskService {
	return &TaskServiceStruct{db: db, ctx: ctx}
}

func (t *TaskServiceStruct) GetAllTasks() ([]models.Task, error) {
	filter := bson.D{{}}
	var tasks []models.Task
	cursor, err := t.db.Find(t.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(t.ctx)
	for cursor.Next(t.ctx) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}
	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	}
	return tasks, nil
}

func (t *TaskServiceStruct) GetSpecificTask(intId primitive.ObjectID) (models.Task, error) {
	filter := bson.D{{
		Key:   "_id",
		Value: intId,
	}}
	var task models.Task
	err := t.db.FindOne(t.ctx, filter).Decode(&task)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (t *TaskServiceStruct) AddNewTask(newTask models.Task) (models.Task, error) {
	added, err := t.db.InsertOne(t.ctx, newTask)
	if err != nil {
		log.Fatal(err)
		return models.Task{}, err
	}
	check, error := t.GetSpecificTask(added.InsertedID.(primitive.ObjectID))
	if error != nil {
		fmt.Print(error)
		return models.Task{}, error
	}
	return check, nil
}

func (t *TaskServiceStruct) DeleteTask(intId primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: intId}}
	delete, err := t.db.DeleteOne(t.ctx, filter)
	if err != nil {
		return err
	}
	if delete.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (t *TaskServiceStruct) UpdateTask(updatedTask models.Task) error {

	taskID := updatedTask.ID
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}
	report, err := t.db.UpdateByID(t.ctx, taskID, update)

	if err != nil {
		return errors.New("failed to update")
	}
	if report.MatchedCount == 0 {
		_, err := t.AddNewTask(updatedTask)
		if err != nil {
			return err
		}

	}
	return nil

}
