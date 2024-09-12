package main

import (
	"fmt"
	"time"
)

func parseSoboceJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Soboce",
		LogoUrl: "https://www.soboce.com/wp-content/uploads/2023/10/Logo-Cabecera.webp",
		JobsUrl: "https://soboce.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
