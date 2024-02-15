package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func parseGrupoVenadoJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Grupo Venado",
		LogoUrl: "https://grupovenado.com/wp-content/uploads/2020/06/Isologo-Venado-Sin-Texto.svg",
		JobsUrl: "https://grupovenado.com/unete-al-equipo-grupo-venado/",
		Jobs:    []Job{},
	}

	c := colly.NewCollector()
	c.OnHTML("div.wpgb-card-media-content-top", func(e *colly.HTMLElement) {
		job := Job{}
		title := e.DOM.Find("a")
		dpto := e.DOM.Find("div:nth-child(2)").Text()
		job.Title = title.Text()
		job.Url = title.AttrOr("href", "")
		job.Depto = parseDepto(dpto)
		job.Area = e.DOM.Find("div:nth-child(3)").Text()
		job.DueDate = e.DOM.Find("div:nth-child(5)").Text()
		parseGrupoVenadoJob(&job, company.Name)
		company.Jobs = append(company.Jobs, job)
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s error: %v\n", company.Name, err)
	})
	c.Visit(company.JobsUrl)

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)
	return company
}
func parseGrupoVenadoJob(job *Job, company string) {
	c := colly.NewCollector()

	c.OnHTML("div.et_pb_row_2_tb_body", func(e *colly.HTMLElement) {
		job.Content, _ = e.DOM.Html()
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s job error (%s): %v\n", company, job.Url, err)
	})

	c.Visit(job.Url)
}
