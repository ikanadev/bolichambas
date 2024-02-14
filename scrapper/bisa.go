package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func ParseBisaJobs() []Job {
	const imgUrl = "https://www.grupobisa.com/assets/IconosEmpresas/BancoBisa.png"
	const company = "Banco Bisa"
	foundJobs := []Job{}
	c := colly.NewCollector(
		colly.AllowedDomains("bancobisa.evaluar.com"),
	)
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
		job.Company = company
		job.ImageUrl = imgUrl
		foundJobs = append(foundJobs, job)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Searching Banco Bisa jobs...")
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error fetching Banco Bisa jobs:", err)
	})
	c.Visit("https://bancobisa.evaluar.com/convocatorias-2/")

	for i := range foundJobs {
		ParseBisaJobDetails(&foundJobs[i])
	}

	return foundJobs
}

func ParseBisaJobDetails(job *Job) {
	c := colly.NewCollector(colly.AllowedDomains("bancobisa.evaluar.com"))
	c.OnHTML("div.job_description", func(e *colly.HTMLElement) {
		job.Content, _ = e.DOM.Html()
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
