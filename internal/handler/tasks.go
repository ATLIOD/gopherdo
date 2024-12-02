package handler

import (
	"encoding/json"
	"gopherdo/internal/model"
	"net/http"
	"strconv"
)

// HandleTasks handles GET and POST requests for tasks
func (h *Handler) handleTasks(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        h.getTasks(w, r)
    case http.MethodPost:
        h.createTask(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// getTasks returns all tasks
func (h *Handler) getTasks(w http.ResponseWriter, r *http.Request) {
    rows, err := h.db.Query(`
        SELECT id, name, description, category, urgency, due_date, complete, created_at, updated_at
        FROM tasks
    `)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var tasks []model.Task
    for rows.Next() {
        var task model.Task
        err := rows.Scan(
            &task.ID,
            &task.Name,
            &task.Description,
            &task.Category,
            &task.Urgency,
            &task.DueDate,
            &task.Complete,
            &task.CreatedAt,
            &task.UpdatedAt,
        )
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tasks = append(tasks, task)
    }

    json.NewEncoder(w).Encode(tasks)
}

// createTask creates a new task
func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
    var task model.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result, err := h.db.Exec(`
        INSERT INTO tasks (name, description, category, urgency, due_date, complete)
        VALUES (?, ?, ?, ?, ?, ?)
    `, task.Name, task.Description, task.Category, task.Urgency, task.DueDate, task.Complete)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, _ := result.LastInsertId()
    task.ID = id

    json.NewEncoder(w).Encode(task)
}

// handleCompleteTask handles task completion
func (h *Handler) handleCompleteTask(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPatch {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    idStr := r.URL.Query().Get("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    _, err = h.db.Exec("UPDATE tasks SET complete = true WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// handleDeleteTask handles task deletion
func (h *Handler) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    idStr := r.URL.Query().Get("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    _, err = h.db.Exec("DELETE FROM tasks WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
