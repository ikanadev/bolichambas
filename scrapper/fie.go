package main

import (
	"fmt"
	"time"
)

func parseFieJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Banco Fie",
		LogoUrl: "https://www.bancofie.com.bo/logo-fie.svg",
		JobsUrl: "https://bancofie.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
