package task

import "time"

// Task represents a task in the ToDo system
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	IsComplete  bool      `json:"is_complete"`
}

// NewTaks creates a new task with ID and timestamp
func NewTask(id int, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		CreatedAt:   time.Now(),
		IsComplete:  false,
	}
}

// Complete marks the task as complete
func (t *Task) Complete() {
	t.IsComplete = true
}
