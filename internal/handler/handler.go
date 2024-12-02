package handler

import (
	"gopherdo/internal/db"
	"net/http"
)

type Handler struct {
	db *db.DB
}

func NewHandler(db *db.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Start(port string) error {
	// Set up your routes here
	http.HandleFunc("/tasks", h.handleTasks)
	http.HandleFunc("/tasks/complete", h.handleCompleteTask)
	http.HandleFunc("/tasks/delete", h.handleDeleteTask)

	// Serve static files
	fs := http.FileServer(http.Dir("api/static"))
	http.Handle("/", fs)

	return http.ListenAndServe(":"+port, nil)
}
