package handler

import (
	"encoding/json"
	"net/http"

	"task-manager/manager"
)

type TaskHandler struct {
	manager *manager.TaskManager
}

func NewTaskHandler(m *manager.TaskManager) *TaskHandler {
	return &TaskHandler{manager: m}
}

type CreateTaskRequest struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	h.manager.AddTask(req.Name, req.Priority)
	w.WriteHeader(http.StatusAccepted)
}
