package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
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
		// case "mark-in-progress":
		// 	handleMarkInProgressCommand(flag.Arg(1))
		// case "mark-done":
		// 	handleMarkDoneCommand(flag.Arg(1))
		// case "list":
		// 	handleListCommand(flag.Arg(1))
	}
}

func handleFileOpening(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open or create file: %w", err)
	}

	stats, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("failed to get file stats: %w", err)
	}

	if stats.Size() == 0 {
		fmt.Println("File is empty or just created. Initializing with '[]'...")

		_, err = file.WriteString("[]")
		if err != nil {
			file.Close()
			return nil, fmt.Errorf("failed to write default JSON: %w", err)
		}

		_, err = file.Seek(0, 0)
		if err != nil {
			file.Close()
			return nil, fmt.Errorf("failed to rewind file cursor: %w", err)
		}
	} else {
		fmt.Println("File exists and already contains data.")
	}

	return file, nil
}

func handleAddCommand(description string, tasks []Task, path string) {

	if description == "" {
		fmt.Println("Please enter the description for the task to be added.")
		return
	}

	now := time.Now().Format(time.DateTime)
	id := len(tasks) + 1
	task := Task{
		Id:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, task)
	writeToFile(tasks, path)
	fmt.Printf("Task added successfully (ID: %d)\n", id)
}

func writeToFile(tasks []Task, path string) {
	jsonBytes, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		return
	}

	err = os.WriteFile(path, jsonBytes, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}

func handleUpdateCommand(id string, description string, tasks []Task, path string) {

	if id == "" || description == "" {
		fmt.Println("Please ensure that both the Id to be updated and new description is present")
		return
	}

	exists, idInt := checkIfIdExists(id, tasks)

	if !exists {
		fmt.Println("Id does not exist")
		return
	}

	index := slices.IndexFunc(tasks, func(t Task) bool {
		return t.Id == idInt
	})

	tasks[index].Description = description

	now := time.Now().Format(time.DateTime)
	tasks[index].UpdatedAt = now

	writeToFile(tasks, path)
}

func handleDeleteCommand(id string, tasks []Task, path string) {

	if id == "" {
		fmt.Println("Please enter the id to be deleted")
		return
	}

	exists, idInt := checkIfIdExists(id, tasks)

	if !exists {
		fmt.Println("Id does not exist")
		return
	}

	tasks = slices.DeleteFunc(tasks, func(t Task) bool {
		return t.Id == idInt
	})

	writeToFile(tasks, path)
}

func checkIfIdExists(id string, tasks []Task) (bool, int) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion:", err)
		return false, -1
	}

	return slices.ContainsFunc(tasks, func(t Task) bool {
		return t.Id == idInt
	}), idInt
}
