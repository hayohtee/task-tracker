package data

// status is a custom type which represents the
// status of a task.
// It uses a string as the underlying type
type status string

const (
	StatusTodo       = status("todo")
	StatusInProgress = status("in-progress")
	StatusDone       = status("done")
)
