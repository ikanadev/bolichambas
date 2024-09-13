package main

import (
	"fmt"
	"time"
)

func parseOpalJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Opal",
		LogoUrl: "https://www.opal.bo/cdn/shop/files/logo_opal.svg",
		JobsUrl: "https://opal.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
