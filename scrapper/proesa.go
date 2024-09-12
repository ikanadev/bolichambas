package main

import (
	"fmt"
	"time"
)

func parseProesaJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Proesa",
		LogoUrl: "http://www.proesabol.com/wp-content/uploads/2021/04/cropped-cropped-marca-proesa.png",
		JobsUrl: "https://proesa.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
