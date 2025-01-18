package data

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

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

// TaskList is a custom type with the underlying data type of
// []task.
// It contains various method for creating and managing tasks.
type TaskList []task

// Get opens the provided file name, retrieve its JSON content and
// decode it into TaskList.
func (t *TaskList) Get(filename string) error {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			return nil
		default:
			return err
		}
	}

	if len(fileContent) == 0 {
		return nil
	}

	return json.Unmarshal(fileContent, t)
}
