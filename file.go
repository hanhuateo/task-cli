package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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
