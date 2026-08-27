package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
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
