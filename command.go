package main

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"time"
)

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
	updateUpdatedAt(tasks, index)
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

	tasks = updateId(tasks)
	writeToFile(tasks, path)
}

func handleMarkInProgressCommand(id string, tasks []Task, path string) {

	if id == "" {
		fmt.Println("Please enter the id to be marked as in progress")
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

	tasks[index].Status = "in-progress"
	updateUpdatedAt(tasks, index)
	writeToFile(tasks, path)
}

func handleMarkDoneCommand(id string, tasks []Task, path string) {
	if id == "" {
		fmt.Println("Please enter the id to be marked as done")
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

	tasks[index].Status = "done"
	updateUpdatedAt(tasks, index)
	writeToFile(tasks, path)
}

func handleListCommand(status string, tasks []Task) {
	for i := range tasks {
		if tasks[i].Status == status || status == "" {
			prettyJSON, err := json.MarshalIndent(tasks[i], "", "	")
			if err != nil {
				fmt.Println("Error generating pretty JSON:", err)
				return
			}
			fmt.Println(string(prettyJSON))
		}
	}
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

func updateUpdatedAt(tasks []Task, index int) {
	now := time.Now().Format(time.DateTime)
	tasks[index].UpdatedAt = now
}

func updateId(tasks []Task) []Task {
	for i := range tasks {
		tasks[i].Id = i + 1
	}
	return tasks
}
