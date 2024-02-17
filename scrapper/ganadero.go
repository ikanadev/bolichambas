package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func parseGanaderoJobs() Company {
	start := time.Now()
	company := Company{
		LogoUrl: "https://www.bg.com.bo/favicon.ico",
		Name:    "Banco Ganadero",
		JobsUrl: "https://www.bg.com.bo/girhr-postulaciones/servlet/hbolsa_trabajo",
		Jobs:    []Job{},
	}
	c := colly.NewCollector()

	// This is crazy, but Banco Ganadero actually saves their data in a hidden input field
	// in JSON format. So we have to parse it manually, and convert it to []Job
	c.OnHTML("input[name='Gr_seleccionesContainerDataV']", func(e *colly.HTMLElement) {
		jsonData := e.Attr("value")
		data := make([][]string, 0)
		json.Unmarshal([]byte(jsonData), &data)
		jobs := make([]Job, len(data))
		baseUrl := "https://www.bg.com.bo/girhr-postulaciones/servlet/hbolsa_trabajo_detalle?"

		for i := range data {
			dates := strings.Split(data[i][33], "-")
			publishDate, _ := time.Parse("02/01/2006", strings.TrimSpace(dates[0]))
			dueDate := strings.TrimSpace(dates[1])
			jobs[i].Depto = parseDepto(data[i][16])
			jobs[i].PublishDate = &publishDate
			jobs[i].DueDate = dueDate
			jobs[i].Title = data[i][61]
			jobs[i].Url = baseUrl + data[i][122]
		}
		company.Jobs = jobs
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s Error: %v\n", company.Name, err)
	})
	c.Visit(company.JobsUrl)

	for i := range company.Jobs {
		parseGanaderoJob(&company.Jobs[i], company.Name)
	}

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)
	return company
}

func parseGanaderoJob(job *Job, company string) {
	c := colly.NewCollector()

	c.OnHTML("div#DL_SELECCIONDETALLE", func(e *colly.HTMLElement) {
		job.Content, _ = e.DOM.Html()
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s job error (%s): %v\n", company, job.Url, err)
	})

	c.Visit(job.Url)
}
