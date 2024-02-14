package main

import "time"

type Company struct {
	Name    string `json:"name"`
	JobsUrl string `json:"jobsUrl"`
	LogoUrl string `json:"logoUrl"`
	Jobs    []Job  `json:"jobs"`
}

// Emty values means that the job does not have that content
type Job struct {
	Title       string     `json:"title"`
	Depto       string     `json:"depto"`
	Url         string     `json:"url"`
	PublishDate *time.Time `json:"publishDate"`
	DueDate     string     `json:"dueDate"`
	Content     string     `json:"content"`
	Area        string     `json:"area"`
}
