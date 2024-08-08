# Task Management API

## The post man documentation can be found at : https://documenter.getpostman.com/view/36144456/2sA3s1nXMh 

## Overview

This API provides endpoints for managing tasks. It supports operations for creating, retrieving, updating, and deleting tasks. Tasks are stored in-memory and each task has a unique ID.

## Endpoints

### 1. Get All Tasks

**GET** `/tasks`

Retrieve all tasks.

**Response:**
- **200 OK**: Returns a list of tasks.
- **404 Not Found**: If there are no tasks.

### 2. Get Specific Task

**GET** `/tasks/:id`

Retrieve a specific task by its ID.

**Path Parameters:**
- `id` (integer): The ID of the task.

**Response:**
- **200 OK**: Returns the requested task.
- **400 Bad Request**: If the ID is invalid.
- **404 Not Found**: If the task with the specified ID does not exist.

### 3. Add New Task

**POST** `/tasks`

Create a new task.

**Request Body:**
- A JSON object representing the task with fields for `Name` and `Description`.

**Response:**
- **201 Created**: Returns the ID of the newly created task.
- **400 Bad Request**: If the request body is invalid.

### 4. Delete Task

**DELETE** `/tasks/:id`

Delete a specific task by its ID.

**Path Parameters:**
- `id` (integer): The ID of the task.

**Response:**
- **200 OK**: If the task was successfully deleted.
- **400 Bad Request**: If the ID is invalid.
- **404 Not Found**: If the task with the specified ID does not exist.

### 5. Update Task

**PUT** `/tasks/:id`

Update an existing task by its ID. If the task does not exist, a new task is created.

**Path Parameters:**
- `id` (integer): The ID of the task.

**Request Body:**
- A JSON object representing the task with fields for `Name` and `Description`.

**Response:**
- **200 OK**: If the task was successfully updated.
- **201 Created**: If a new task was created because the specified ID did not exist.
- **400 Bad Request**: If the request body is invalid.

## Error Handling

- **400 Bad Request**: Invalid input, such as invalid task data or ID format.
- **404 Not Found**: The requested resource does not exist.

## Notes

- The task ID is auto-incremented and starts from 1.
- All tasks are stored in-memory, so they will be lost if the server is restarted.
