package main

import (
	"fmt"
	"time"
)

func ParseSolJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Banco Sol",
		LogoUrl: "https://cdn.modyo.cloud/uploads/58dd57da-031f-42e8-ba3c-c5d78bffc7fb/original/logo_bancosol.svg",
		JobsUrl: "https://bancosol.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
