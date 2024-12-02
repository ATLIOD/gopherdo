package model

import "time"

type Task struct {
	ID          int64      `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	Category    string     `json:"category" db:"category"`
	Urgency     string     `json:"urgency" db:"urgency"`
	DueDate     *time.Time `json:"dueDate" db:"due_date"`
	Complete    bool       `json:"complete" db:"complete"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time  `json:"updatedAt" db:"updated_at"`
}
