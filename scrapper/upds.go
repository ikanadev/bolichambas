package main

import (
	"fmt"
	"time"
)

func parseUPDSJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Universidad privada Domingo Savio",
		LogoUrl: "https://www.upds.edu.bo/wp-content/uploads/2020/10/upds_logo-1-1-1.png",
		JobsUrl: "https://upds.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
