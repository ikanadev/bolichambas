package main

import (
	"fmt"
	"time"
)

func parseBdpJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Banco de Desarrollo Productivo",
		LogoUrl: "https://www.bdp.com.bo/wp-content/uploads/2022/04/BDP_Logo-SVG.svg",
		JobsUrl: "https://bdp.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
