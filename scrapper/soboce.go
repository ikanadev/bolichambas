package main

import (
	"fmt"
	"time"
)

func parseSoboceJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Soboce",
		LogoUrl: "https://soboce.com/wp-content/uploads/2025/02/Logo-Soboce.webp",
		JobsUrl: "https://soboce.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
