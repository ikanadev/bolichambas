package main

import (
	"fmt"
	"time"
)

func ParseDiaconiaJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Diaconía",
		LogoUrl: "https://www.diaconia.bo/web/images/logo.svg",
		JobsUrl: "https://diaconia.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
