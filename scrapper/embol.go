package main

import (
	"fmt"
	"time"
)

func ParseEmbolJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Embol",
		LogoUrl: "https://embol.com/small/embolHeaderIconDark.svg",
		JobsUrl: "https://embol.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
