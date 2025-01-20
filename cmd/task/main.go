package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/hayohtee/task-tracker/internal/data"
)

// taskFileName is the default file name for the task-cli.
var taskFileName = ".task.json"

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	markInProgressCmd := flag.NewFlagSet("mark-in-progress", flag.ExitOnError)
	markDoneCmd := flag.NewFlagSet("mark-done", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Check if a sub-command is not provided and exit with an error.
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no sub-commands provided")
		os.Exit(1)
	}

	// Check if the TASK_FILENAME environment variable is set and update the taskFileName.
	if os.Getenv("TASK_FILENAME") != "" {
		taskFileName = os.Getenv("TASK_FILENAME")
	}

	// Declare a TaskList variable and attempt to read the contents from the filename.
	var taskList data.TaskList
	if err := taskList.Get(taskFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		args := addCmd.Args()
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "required an argument containing the name of the task")
			os.Exit(1)
		}
		taskList.Add(args[0])

		// Save the new list.
		if err := taskList.Save(taskFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Print successful message
		fmt.Printf("Task added successfully (ID: %d)\n", taskList[len(taskList)-1].ID)

	case "update":
		updateCmd.Parse(os.Args[2:])
		args := updateCmd.Args()
		if len(args) != 2 {
			fmt.Fprintln(os.Stderr, "required 2 arguments and in the format <<1 \"New Task Name\">>")
			os.Exit(1)
		}

		// Convert the first argument to an int
		pos, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Update the task
		if err := taskList.Update(pos, args[1]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list.
		if err := taskList.Save(taskFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		args := deleteCmd.Args()
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "required an argument containing the position of the task")
			os.Exit(1)
		}

		// Convert the first argument to an int.
		pos, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Delete the task.
		if err := taskList.Delete(pos); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list.
		if err := taskList.Save(taskFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
