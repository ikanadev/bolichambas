package main

import (
	"fmt"
	"time"
)

func ParseBNBJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Banco Nacional de Bolivia",
		LogoUrl: "https://www.bnb.com.bo/PortalBNB/Images/PNG_BNB_Blanco.png",
		JobsUrl: "https://bnb.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}
	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
