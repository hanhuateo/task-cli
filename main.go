package main

import (
	"encoding/json"
	"flag"
	"log"
)

func main() {
	command := flag.String("command", "", "The command to execute")
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatalf("There is no command to execute. %s", *command)
	}

	path := "tasks.json"
	file, err := handleFileOpening(path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var tasks []Task

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	switch flag.Arg(0) {
	case "add":
		handleAddCommand(flag.Arg(1), tasks, path)
	case "update":
		handleUpdateCommand(flag.Arg(1), flag.Arg(2), tasks, path)
	case "delete":
		handleDeleteCommand(flag.Arg(1), tasks, path)
	case "mark-in-progress":
		handleMarkInProgressCommand(flag.Arg(1), tasks, path)
	case "mark-done":
		handleMarkDoneCommand(flag.Arg(1), tasks, path)
	case "list":
		handleListCommand(flag.Arg(1), tasks)
	}
}
