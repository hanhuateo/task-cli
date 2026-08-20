package main

import "time"

type Task struct {
	id          string
	description string
	status      string
	createdAt   time.Time
	updatedAt   time.Time
}
