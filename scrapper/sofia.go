package main

import (
	"fmt"
	"time"
)

func ParseSofiaJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Sofía",
		LogoUrl: "https://sofia.com.bo/wp-content/themes/sofia/images/sofia-logo.svg",
		JobsUrl: "https://sofia.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
