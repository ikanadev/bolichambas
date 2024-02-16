package main

import (
	"fmt"
	"time"
)

func parseMonopolJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Monopol",
		LogoUrl: "https://www.pinturasmonopol.com/wp-content/uploads/2024/01/monopol_.svg",
		JobsUrl: "https://monopol.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
