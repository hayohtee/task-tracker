package data

import "time"

// status is a custom type which represents the
// status of a task.
// It uses a string as the underlying type
type status string

const (
	StatusTodo       = status("todo")
	StatusInProgress = status("in-progress")
	StatusDone       = status("done")
)

// task is a struct that holds all the information
// about a particular task.
type task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
