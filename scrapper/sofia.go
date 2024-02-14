package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func ParseSofiaJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Sofía",
		LogoUrl: "https://sofia.com.bo/wp-content/themes/sofia/images/sofia-logo.svg",
		JobsUrl: "https://sofia.evaluar.com/",
		Jobs:    []Job{},
	}
	c := colly.NewCollector()

	c.OnHTML("a.job_listing", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		info := e.DOM.ChildrenFiltered("div.listing-title")
		title := strings.ReplaceAll(info.ChildrenFiltered("h4").Text(), "\n", "")
		depto := info.ChildrenFiltered("ul").ChildrenFiltered("li:nth-child(2)").Text()
		depto = parseDepto(depto)
		job := Job{}
		job.Url = link
		job.Title = title
		job.Depto = depto
		company.Jobs = append(company.Jobs, job)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s error: %v\n", company.Name, err)
	})

	c.Visit(company.JobsUrl)

	for i := range company.Jobs {
		ParseSofiaJobDetails(&company.Jobs[i], company.Name)
	}
	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}

func ParseSofiaJobDetails(job *Job, company string) {
	c := colly.NewCollector()
	c.OnHTML("div.job_description", func(e *colly.HTMLElement) {
		job.Content, _ = e.DOM.Html()
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s job error (%s): %v\n", company, job.Url, err)
	})
	c.OnHTML("div.job-overview", func(e *colly.HTMLElement) {
		publishDateStr, _ := e.DOM.Find("time").Attr("datetime")
		dueDate := e.DOM.ChildrenFiltered("ul").ChildrenFiltered("li:nth-child(2)").Find("span").Text()
		publishDate, _ := time.Parse("2006-01-02", publishDateStr)
		job.PublishDate = &publishDate
		job.DueDate = dueDate
	})
	c.Visit(job.Url)
}
