package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "task-cli"
	fileName = ".task.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	// Append .exe extension to binName if the operating
	// system that is running the test is windows.
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot build tool %s: %s", binName, err.Error())
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTaskCLI(t *testing.T) {
	tasks := []string{"Test Task 1", "Test Task 2", "Test Task 3", "Test Task 4", "Test Task 5"}

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("Add new tasks", func(t *testing.T) {
		for index, task := range tasks {
			cmd := exec.Command(cmdPath, "add", task)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}

			expected := fmt.Sprintf("Task added successfully (ID: %d)\n", index+1)
			if string(out) != expected {
				t.Errorf("expected %q but got %q instead\n", expected, string(out))
			}
		}
	})

	t.Run("Update Some Tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "update", "1", "New Test Task 1")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		cmd = exec.Command(cmdPath, "update", "3", "New Test Task 3")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Delete Some Tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "delete", "5")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		cmd = exec.Command(cmdPath, "delete", "4")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Mark A Task As In-Progress", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "mark-in-progress", "2")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Mark A Task As Done", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "mark-done", "3")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("List All Tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("%d. %s\tstatus: %q\n", 1, "New Test Task 1", "todo")
		expected += fmt.Sprintf("%d. %s\tstatus: %q\n", 2, tasks[1], "in-progress")
		expected += fmt.Sprintf("%d. %s\tstatus: %q\n\n", 3, "New Test Task 3", "done")

		if string(out) != expected {
			t.Errorf("expected %q but got %q instead", expected, string(out))
		}
	})

	t.Run("List Todo Tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "list", "todo")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("%d. %s\n", 1, "New Test Task 1")
		if string(out) != expected {
			t.Errorf("expected %q but got %q instead", expected, string(out))
		}
	})

	t.Run("List Done Tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "list", "done")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("%d. %s\n", 1, "New Test Task 3")
		if string(out) != expected {
			t.Errorf("expected %q but got %q instead", expected, string(out))
		}
	})
}
