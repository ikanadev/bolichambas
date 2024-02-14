package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func saveData(companies *[]Company) {
	now := time.Now()
	jsonData, err := json.MarshalIndent(companies, "", "  ")
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("jobs_%s_farmacorp.json", now.Format("2006-01-02"))
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	companies := []Company{}
	// companies = append(companies, ParseBisaJobs())
	companies = append(companies, ParseFarmacorpJobs())
	saveData(&companies)
}
