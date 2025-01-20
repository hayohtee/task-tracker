package main

import (
	"flag"
	"fmt"
	"os"
)

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

}
