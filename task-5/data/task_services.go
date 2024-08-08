package data

import (
	"backend-learning-track/task-4/models"
	"errors"
)



var tasks []models.Task

func GetAllTasks() ([]models.Task, error) {
	if len(tasks) == 0 {
		return nil, errors.New("no tasks found")
	}
	return tasks, nil
}

func GetSpecificTask(intId int) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == intId {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func AddNewTask(newTask models.Task) {
	tasks = append(tasks, newTask)
}

func DeleteTask(intId int) error {
	for i, task := range tasks {
		if task.ID == intId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")

}

func UpdateTask(updatedTask models.Task) (int, error) {
	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			tasks[i] = updatedTask
			return 204, nil
		}
	}
	tasks = append(tasks, updatedTask)
	return 201, nil
}
