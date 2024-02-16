package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func parseDatecJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Datec",
		LogoUrl: "https://www.datec.com.bo/wp-content/uploads/2023/06/logo-datec.png",
		JobsUrl: "https://www.datec.com.bo/postulaciones/",
		Jobs:    []Job{},
	}

	c := colly.NewCollector()
	c.OnHTML("div.jet-listing-grid__item", func(e *colly.HTMLElement) {
		title := e.DOM.Find("h4").Text()
		url := e.DOM.Find("a").AttrOr("href", "")
		company.Jobs = append(company.Jobs, Job{Title: title, Url: url})
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s error: %v\n", company.Name, err)
	})
	c.Visit(company.JobsUrl)

	for i := range company.Jobs {
		parseDatecJob(&company.Jobs[i], company.Name)
	}

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)

	return company
}
func parseDatecJob(job *Job, company string) {
	body, err := fetchUrl(job.Url)
	if err != nil {
		fmt.Printf("%s job error(%s): %v\n", company, job.Url, err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	sections := doc.Find("section.elementor-section")
	depto := parseDepto(sections.Eq(0).Find("li").Text())
	content, _ := sections.Eq(1).Children().Html()
	job.Depto = depto
	job.Content = content
}
