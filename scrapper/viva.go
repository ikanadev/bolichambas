package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func parseVivaJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Viva",
		LogoUrl: "https://www.viva.com.bo/wp-content/uploads/2022/09/Logo-Viva.png",
		JobsUrl: "https://talentoviva.hr-suite.app/vacantes.php",
		Jobs:    []Job{},
	}

	c := colly.NewCollector()

	c.OnHTML("div.card", func(e *colly.HTMLElement) {
		title := e.DOM.Find("h2").Text()
		url := e.DOM.Find("a").AttrOr("href", "")
		company.Jobs = append(company.Jobs, Job{Title: title, Url: url})
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s error: %v\n", company.Name, err)
	})
	c.Visit(company.JobsUrl)

	for i := range company.Jobs {
		parseVivaJob(&company.Jobs[i], company.Name)
	}

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)
	return company
}
func parseVivaJob(job *Job, company string) {
	c := colly.NewCollector()
	c.OnHTML("div.card-body", func(e *colly.HTMLElement) {
		dataCont := e.DOM.Find("p").First()
		content := dataCont.Find("b")
		publishDate, _ := time.Parse("2006-01-02", content.Eq(3).Text()[0:10])
		job.Depto = parseDepto(content.Eq(0).Text())
		job.PublishDate = &publishDate
		job.DueDate = content.Eq(5).Text()
		job.Area = content.Eq(6).Text()

		desc := e.DOM.Find("div").Eq(1).Slice(0, 3)
		html, _ := desc.Html()
		job.Content = html
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s job error(%s): %v\n", company, job.Url, err)
	})
	c.Visit(job.Url)
}
