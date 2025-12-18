package main

import (
	"log"
	"net/http"

	"task-manager/handler"
	"task-manager/manager"
)

func main() {
	taskManager := manager.NewTaskManager(3)
	taskHandler := handler.NewTaskHandler(taskManager)

	http.HandleFunc("/tasks", taskHandler.CreateTask)

	log.Println("Task Manager running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
