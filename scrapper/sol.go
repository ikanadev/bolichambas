package main

import (
	"fmt"
	"time"
)

func ParseSolJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Banco Sol",
		LogoUrl: "https://www.bancosol.com.bo/wp-content/uploads/2022/07/Logo-.png",
		JobsUrl: "https://bancosol.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
