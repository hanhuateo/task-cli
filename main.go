package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	command := flag.String("command", "", "The command to execute")
	flag.Parse()

	// check for arguments first
	if flag.NArg() == 0 {
		log.Fatalf("there is no command : %s", *command)
	}

	path := "tasks.json"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	if errors.Is(err, os.ErrExist) {
		fmt.Println("File already exists. Opening...")
		file, err = os.OpenFile(path, os.O_RDWR, 0666)
		if err != nil {
			log.Fatalf("Failed to open existing file: %v", err)
		}
	} else if err != nil {
		log.Fatalf("System error: %v", err)
	} else {
		fmt.Println("File did not exist. Created new file and writing empty JSON structure...")
		_, err = file.WriteString("{}")
		if err != nil {
			log.Fatalf("Failed to write initial JSON data: %v", err)
		}
		file.Seek(0, 0)
	}
	defer file.Close()

	// read json file first
	// fileBytes, err := os.ReadFile("tasks.json")
	// if err != nil {
	// 	log.Fatalf("Error reading file: %v", err)
	// }

	// fmt.Printf("fileBytes : %s", fileBytes)
	// fmt.Printf("length : %d", len(fileBytes))

	// var tasks Task

	// err = json.Unmarshal(fileBytes, &tasks)
	// if err != nil {
	// 	log.Fatalf("Error parsing JSON: &v", err)
	// }

	// switch flag.Arg(0) {
	// case "add":
	// 	fmt.Println("add")
	// case "update":
	// 	fmt.Println("update")
	// case "delete":
	// 	fmt.Println("delete")
	// case "mark-in-progress":
	// 	fmt.Println("mark-in-progress")
	// case "mark-done":
	// 	fmt.Println("mark-done")
	// case "list":
	// 	fmt.Println("list")
	// }
}

func taskFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}
