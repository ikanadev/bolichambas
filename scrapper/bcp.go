package main

import (
	"fmt"
	"time"
)

func parseBcpJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Banco de Crédito Bolivia",
		LogoUrl: "https://www.bcp.com.bo/Content/images/principal/Logo_BCP.svg",
		JobsUrl: "https://bcpbolivia.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
