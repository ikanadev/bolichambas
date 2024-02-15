package main

import (
	"fmt"
	"time"
)

func parseLaPapeleraJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "La Papelera",
		LogoUrl: "https://cdn.evaluar.com/jobboards/prod/wp-content/uploads/jobboards/sites/313/2021/05/04154619/ALE-LPSA.jpg",
		JobsUrl: "https://lapapelera.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
