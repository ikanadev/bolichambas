package main

import (
	"fmt"
	"time"
)

func parseUnionAgroJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Unión Agronegocios",
		LogoUrl: "https://www.union.com.bo/wp-content/uploads/2021/12/logo-principal.svg",
		JobsUrl: "https://unionagronegocios.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
