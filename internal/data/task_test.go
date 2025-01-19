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

func TestAdd(t *testing.T) {
	var taskList data.TaskList

	// Add new tasks to the list
	tasks := []string{"Test Task 1", "Test Task 2"}
	for _, task := range tasks {
		taskList.Add(task)
	}

	// Check if the length are the same
	if len(taskList) != len(tasks) {
		t.Errorf("expected length %d but got %d instead", 2, len(taskList))
	}

	// Check if the tasks are added in ordered manner
	if taskList[0].Description != tasks[0] {
		t.Errorf("expected %q but got %q instead", tasks[0], taskList[0].Description)
	}

	if taskList[1].Description != tasks[1] {
		t.Errorf("expected %q but got %q instead", tasks[1], taskList[1].Description)
	}

	// Check if the ID of the added tasks are incremented
	if taskList[0].ID != 1 {
		t.Errorf("expect the ID of the first task to be %d but got %d", 1, taskList[0].ID)
	}

	if taskList[1].ID != 2 {
		t.Errorf("expect the ID of the second task to be %d but got %d", 2, taskList[0].ID)
	}
}
