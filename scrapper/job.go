package main

import "time"

// Emty values means that the job does not have that content
type Job struct {
	Title       string     `json:"title"`
	ImageUrl    string     `json:"imageUrl"`
	Depto       string     `json:"depto"`
	Company     string     `json:"company"`
	Url         string     `json:"url"`
	PublishDate *time.Time `json:"publishDate"`
	DueDate     string     `json:"dueDate"`
	Content     string     `json:"content"`
}
