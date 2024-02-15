package main

import (
	"fmt"
	"time"
)

func parseBagoJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Laboratorios Bagó",
		LogoUrl: "https://www.bago.com.bo/wp-content/uploads/2022/06/bago.png",
		JobsUrl: "https://bagobolivia.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
