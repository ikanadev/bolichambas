package main

import (
	"fmt"
	"time"
)

func parseBisaJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Banco Bisa",
		LogoUrl: "https://www.grupobisa.com/assets/IconosEmpresas/BancoBisa.png",
		JobsUrl: "https://bancobisa.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}
	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
