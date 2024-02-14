package main

import (
	"encoding/json"
	"os"
)

func saveJobs(jobs *[]Job) {
	jsonData, err := json.MarshalIndent(jobs, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("jobs.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	jobs := []Job{}
	bisaJobs :=ParseBisaJobs()
	jobs = append(jobs, bisaJobs...)
	saveJobs(&jobs)
}
