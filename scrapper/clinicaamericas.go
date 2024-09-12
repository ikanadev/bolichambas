package main

import (
	"fmt"
	"time"
)

func parseClinicaAmericasJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Clínica de las Américas",
		LogoUrl: "https://www.clinicadelasamericas.com.bo/wp-content/uploads/2022/05/clinica-de-las-americas-logo-white.png",
		JobsUrl: "https://clinicadelasamericas.evaluar.com/jm-ajax/get_listings/?per_page=100",
		Jobs:    []Job{},
	}

	parseEvaluar(&company)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
