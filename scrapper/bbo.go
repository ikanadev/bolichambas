package main

import (
	"fmt"
	"time"
)

func parseBBOJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Bebidas Bolivianas",
		LogoUrl: "https://www.bebidasbolivianas.com/wp-content/uploads/2020/08/LogoBBO.png",
		JobsUrl: "https://bbo.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
