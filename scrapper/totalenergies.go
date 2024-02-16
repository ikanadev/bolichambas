package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func parseTotalEnergiesJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Total Energies",
		LogoUrl: "https://totalenergies.com/themes/custom/totalenergies_com/dist/img/logo_totalenergies.png",
		JobsUrl: "https://totalenergies.avature.net/es_ES/careers/SearchJobs/?3834=%5B41569%5D&3834_format=3639&listFilterMode=1&jobRecordsPerPage=6&",
		Jobs:    []Job{},
	}

	c := colly.NewCollector()

	c.OnHTML("h3.article__header__text__title", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(strings.ReplaceAll(e.Text, "\n", " "))
		url := e.DOM.Find("a").First().AttrOr("href", "")
		company.Jobs = append(company.Jobs, Job{Title: title, Url: url})
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s error: %v\n", company.Name, err)
	})
	c.Visit(company.JobsUrl)

	for i := range company.Jobs {
		parseTotalEnergiesJob(&company.Jobs[i], company.Name)
	}

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)
	return company
}
func parseTotalEnergiesJob(job *Job, company string) {
	c := colly.NewCollector()

	section := 1
	c.OnHTML("div.section__content", func(e *colly.HTMLElement) {
		if section > 2 {
			return
		}
		if section == 1 {
			depto := e.DOM.Find("div.article__content__view__field__value").Eq(1).Text()
			depto = parseDepto(depto)
			job.Depto = depto
		}
		if section == 2 {
			job.Content, _ = e.DOM.Html()
		}
		section++
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s job error(%s): %v\n", company, job.Url, err)
	})
	c.Visit(job.Url)
}
