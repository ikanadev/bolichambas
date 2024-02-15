package main

import (
	"fmt"
	"time"
)

func parseAlianzaJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Alianza Seguros",
		LogoUrl: "https://www1.alianza.com.bo/Content/img/logo.png",
		JobsUrl: "https://alianzaseguros.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
