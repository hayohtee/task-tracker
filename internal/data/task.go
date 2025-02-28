package data

import (
	"encoding/json"
	"errors"
	"fmt"
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

// Save converts the TaskList into JSON format and
// writes the content to a file using the specified name.
func (t *TaskList) Save(filename string) error {
	js, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

// Add insert a new task into TaskList.
func (t *TaskList) Add(description string) {
	taskList := *t
	var id int

	if len(taskList) == 0 {
		id = 1
	} else {
		id = taskList[len(taskList)-1].ID + 1
	}

	item := task{
		ID:          id,
		Description: description,
		CreatedAt:   time.Now(),
		Status:      StatusTodo,
	}

	*t = append(*t, item)
}

// Update the description based on the value provided and
// set the updated_at to the current timestamp.
func (t *TaskList) Update(pos int, description string) error {
	if pos <= 0 || pos > len(*t) {
		return fmt.Errorf("task %d does not exist", pos)
	}

	taskList := *t
	taskList[pos-1].UpdatedAt = time.Now()
	taskList[pos-1].Description = description

	return nil
}

// Delete a task from the task list based on its position.
func (t *TaskList) Delete(pos int) error {
	if pos <= 0 || pos > len(*t) {
		return fmt.Errorf("task %d does not exist", pos)
	}

	taskList := *t
	*t = append(taskList[:pos-1], taskList[pos:]...)

	return nil
}

// Mark update the status of a task and also set the
// CreatedAt field to the current timestamp.
func (t *TaskList) Mark(pos int, status status) error {
	if pos <= 0 || pos > len(*t) {
		return fmt.Errorf("task %d does not exist", pos)
	}

	taskList := *t
	taskList[pos-1].Status = status
	taskList[pos-1].UpdatedAt = time.Now()

	return nil
}

// List returns a string containing the list of all task in the TaskList.
func (t *TaskList) List() string {
	total := ""
	for index, value := range *t {
		total += fmt.Sprintf("%d. %s\tstatus: %q\n", index+1, value.Description, value.Status)
	}
	return total
}

// ListByStatus returns a string containing the list of all tasks in the TaskList
// that match the given status.
func (t *TaskList) ListByStatus(status status) string {
	total := ""
	counter := 0
	for _, value := range *t {
		if value.Status == status {
			counter++
			total += fmt.Sprintf("%d. %s\n", counter, value.Description)
		}
	}
	return total
}
