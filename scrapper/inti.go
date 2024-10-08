package main

import (
	"fmt"
	"time"
)

func parseIntiJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "INTI",
		LogoUrl: "https://www.inti.com.bo/wp-content/uploads/2023/07/Logo-INTI-Salud.png",
		JobsUrl: "https://talento.inti.com.bo/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
