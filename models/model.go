package models

import "time"

type Item struct {
	Id        int       `json:"id,omitempty"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	DueDate   time.Time `json:"dueDate"`
}
