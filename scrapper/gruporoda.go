package main

import (
	"fmt"
	"time"
)

func parseGrupoRodaJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Grupo Roda",
		LogoUrl: "https://cdn.evaluar.com/jobboards/prod/wp-content/uploads/jobboards/sites/91/2019/04/06200650/grupo-roda-logo-blanco-1.jpg",
		JobsUrl: "https://gruporoda.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
