package main

import (
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
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Failed to open or create file: %v", err)
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to get file stats: %v", err)
	}

	if stats.Size() == 0 {
		fmt.Println("File is empty or just created. Initializing with '{}'...")

		_, err = file.WriteString("{}")
		if err != nil {
			log.Fatalf("Failed to write default JSON: %v", err)
		}

		_, err = file.Seek(0, 0)
		if err != nil {
			log.Fatalf("Failed to rewind file cursor: %v", err)
		}
	} else {
		fmt.Println("File exists and already contains data.")
	}
}
