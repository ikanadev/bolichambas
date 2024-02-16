package main

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func parseUnividaJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Univida Seguros",
		LogoUrl: "https://www.pinturasmonopol.com/wp-content/uploads/2024/01/monopol_.svg",
		JobsUrl: "https://www.univida.bo/index.php?mod=rrhh&view=announcementspublic",
		Jobs:    []Job{},
	}

	c := colly.NewCollector()
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.DOM.Find("tr").Each(func(i int, tr *goquery.Selection) {
			depto := parseDepto(tr.Find("td").Eq(1).Text())
			title := tr.Find("td").Eq(2).Text()
			dueDate := tr.Find("td").Eq(3).Text()
			url := tr.Find("td").Eq(4).Find("a").AttrOr("href", "")
			company.Jobs = append(
				company.Jobs,
				Job{Title: title, Url: url, DueDate: dueDate, Depto: depto},
			)
		})
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s error: %v\n", company.Name, err)
	})
	c.Visit(company.JobsUrl)

	for i := range company.Jobs {
		parseUnividaJob(&company.Jobs[i], company.Name)
	}

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)
	return company
}
func parseUnividaJob(job *Job, company string) {
	c := colly.NewCollector()
	c.OnHTML("div#announcement-content", func(e *colly.HTMLElement) {
		job.Content, _ = e.DOM.Html()
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s job error(%s): %v\n", company, job.Url, err)
	})
	c.Visit(job.Url)
}
