package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func parseCamsaJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Camsa",
		LogoUrl: "https://camsacorp.bo/wp-content/uploads/2021/09/LOGO-CAMSA-AZUL.png",
		JobsUrl: "https://camsacorp.bo/oportunidades-laborales/",
		Jobs:    []Job{},
	}

	c := colly.NewCollector()
	c.OnHTML("div.job-preview", func(e *colly.HTMLElement) {
		title := e.DOM.Find("h5").Text()
		url := e.DOM.Find("a").Eq(0).AttrOr("href", "")
		depto := parseDepto(e.DOM.Find("span").Eq(0).Text())
		company.Jobs = append(company.Jobs, Job{Title: title, Url: url, Depto: depto})
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s error: %v\n", company.Name, err)
	})
	c.Visit(company.JobsUrl)

	for i := range company.Jobs {
		parseCamsaJob(&company.Jobs[i], company.Name)
	}

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)
	return company
}
func parseCamsaJob(job *Job, company string) {
	body, err := fetchUrl(job.Url)
	if err != nil {
		fmt.Printf("%s job error(%s): %v\n", company, job.Url, err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("%s job error(%s): %v\n", company, job.Url, err)
	}
	content, _ := doc.Find("div.job-content").Html()
	job.Content = content
	job.DueDate = doc.Find("div.type-date-posted").First().Find("div.jobs-row-input").Text()
	job.DueDate = strings.TrimSpace(job.DueDate)
}
