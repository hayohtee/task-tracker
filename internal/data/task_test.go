package data_test

import (
	"os"
	"testing"

	"github.com/hayohtee/task-tracker/internal/data"
)

func TestSaveAndGet(t *testing.T) {
	// Create a temp file for saving and retrieving the task.
	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	var taskList data.TaskList

	// Add a task to TaskList.
	task := "Test Task 1"
	taskList.Add(task)

	// Save the new task to the temp file.
	if err := taskList.Save(tempFile.Name()); err != nil {
		t.Fatal(err)
	}

	var taskList2 data.TaskList

	// Retrieve the saved tasks from temp file.
	if err := taskList2.Get(tempFile.Name()); err != nil {
		t.Fatal(err)
	}

	// Check if the saved task was present in the retrieved task
	if taskList2[0].Description != task {
		t.Errorf("expected %q, got %q instead", task, taskList2[0].Description)
	}
}
