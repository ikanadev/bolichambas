package main

import (
	"fmt"
	"time"
)

func parsePilAndinaJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Pil Andina",
		LogoUrl: "https://pilandina.com.bo/wp-content/uploads/2023/04/Logo-PIL.svg",
		JobsUrl: "https://pilandina.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
