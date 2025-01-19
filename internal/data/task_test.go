package data_test

import (
	"fmt"
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

func TestUpdate(t *testing.T) {
	var taskList data.TaskList

	// Add a new task.
	task := "Test Task"
	taskList.Add(task)

	// Update the added task.
	newTask := "Test New Task"
	if err := taskList.Update(1, newTask); err != nil {
		t.Fatal(err)
	}

	// Check if the task was updated successfully.
	if taskList[0].Description != newTask {
		t.Errorf("expected %q but got %q instead", newTask, taskList[0].Description)
	}
}

func TestDelete(t *testing.T) {
	var taskList data.TaskList

	// Add new tasks to the list.
	tasks := []string{"Test Task 1", "Test Task 2"}
	for _, task := range tasks {
		taskList.Add(task)
	}

	// Check if the length of the list is the same initially.
	if len(taskList) != len(tasks) {
		t.Errorf("expected length %d but got %d instead", 2, len(taskList))
	}

	// Delete the second task
	if err := taskList.Delete(2); err != nil {
		t.Fatal(err)
	}

	// Check if the length of the task list is 1
	if len(taskList) != 1 {
		t.Error("task list should contain only one task")
	}

	// Check if the remaining task is equal to the first task
	if taskList[0].Description != tasks[0] {
		t.Errorf("expected %q but got %q instead", tasks[0], taskList[0].Description)
	}
}

func TestMark(t *testing.T) {
	var taskList data.TaskList

	// Add new tasks to the list.
	tasks := []string{"Test Task 1", "Test Task 2"}
	for _, task := range tasks {
		taskList.Add(task)
	}

	// Mark the first task as in-progress
	if err := taskList.Mark(1, data.StatusInProgress); err != nil {
		t.Fatal(err)
	}

	// Check if the status of the first task is updated successfully.
	if taskList[0].Status != data.StatusInProgress {
		t.Errorf("expected status %q but got %q instead", data.StatusInProgress, taskList[0].Status)
	}

	// Mark the second task as done.
	if err := taskList.Mark(2, data.StatusDone); err != nil {
		t.Fatal(err)
	}

	// Check if the status of the second task is updated successfully.
	if taskList[1].Status != data.StatusDone {
		t.Errorf("expected status %q but got %q instead", data.StatusDone, taskList[1].Status)
	}
}

func TestList(t *testing.T) {
	var taskList data.TaskList

	// Add new tasks to the list.
	tasks := []string{"Test Task 1", "Test Task 2"}
	for _, task := range tasks {
		taskList.Add(task)
	}

	expected := fmt.Sprintf("%d. %s\tstatus: %q\n", 1, tasks[0], data.StatusTodo)
	expected += fmt.Sprintf("%d. %s\tstatus: %q\n", 2, tasks[1], data.StatusTodo)

	got := taskList.List()

	if got != expected {
		t.Errorf("expected %q but got %q instead", expected, got)
	}
}

func TestListByStatus(t *testing.T) {
	var taskList data.TaskList

	// Add new tasks to the list.
	tasks := []string{"Test Task 1", "Test Task 2", "Test Task 3", "Test Task 4", "Test Task 5"}
	for _, task := range tasks {
		taskList.Add(task)
	}

	// Mark task 2 and task 3 as done
	if err := taskList.Mark(2, data.StatusDone); err != nil {
		t.Fatal(err)
	}

	if err := taskList.Mark(3, data.StatusDone); err != nil {
		t.Fatal(err)
	}

	expected := fmt.Sprintf("%d. %s\n", 1, tasks[1])
	expected += fmt.Sprintf("%d. %s\n", 2, tasks[2])

	got := taskList.ListByStatus(data.StatusDone)

	// Check if it list all done tasks correctly.
	if got != expected {
		t.Errorf("expected %q but got %q instead", expected, got)
	}

	// Mark task 1, 4 and 5 as in-progress
	if err := taskList.Mark(1, data.StatusInProgress); err != nil {
		t.Fatal(err)
	}
	if err := taskList.Mark(4, data.StatusInProgress); err != nil {
		t.Fatal(err)
	}
	if err := taskList.Mark(5, data.StatusInProgress); err != nil {
		t.Fatal(err)
	}

	expected = fmt.Sprintf("%d. %s\n", 1, tasks[0])
	expected += fmt.Sprintf("%d. %s\n", 2, tasks[3])
	expected += fmt.Sprintf("%d. %s\n", 3, tasks[4])

	got = taskList.ListByStatus(data.StatusInProgress)

	// Check if it list all done tasks correctly.
	if got != expected {
		t.Errorf("expected %q but got %q instead", expected, got)
	}
}
