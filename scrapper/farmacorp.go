package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gocolly/colly"
)

func ParseFarmacorpJobs() Company {
	type Data struct {
		Data struct {
			Offers []struct {
				Name            string `json:"name"`
				Region          string `json:"region"`
				Link            string `json:"link"`
				PublicationDays int    `json:"publication_days"`
				Area            string `json:"area"`
			} `json:"offers"`
		} `json:"data"`
	}

	start := time.Now()
	company := Company{
		Name:    "Farmacorp",
		LogoUrl: "https://farmacorp.com/cdn/shop/files/farmacorp_-_copia_672x180_671x180_671x180_1570e736-2830-4a88-9469-2cffd4cb908e_671x180.webp",
		JobsUrl: "https://gcs-storage.airavirtual.com/public/feeds/aira_farmacorp.json",
		Jobs:    []Job{},
	}
	response, err := http.Get(company.JobsUrl)
	if err != nil {
		fmt.Printf("%s Error: %v\n", company.Name, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s Error: %v\n", company.Name, err)
	}

	data := Data{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("%s Error: %v\n", company.Name, err)
	}
	jobs := make([]Job, len(data.Data.Offers))
	now := time.Now()
	for i := range data.Data.Offers {
		publishDate := now.AddDate(0, 0, -data.Data.Offers[i].PublicationDays)
		jobs[i].Title = data.Data.Offers[i].Name
		jobs[i].Url = data.Data.Offers[i].Link
		jobs[i].Area = data.Data.Offers[i].Area
		jobs[i].Depto = parseDepto(data.Data.Offers[i].Region)
		jobs[i].PublishDate = &publishDate
	}
	company.Jobs = jobs

	for i := range company.Jobs {
		parseFarmacorpJobContent(&company.Jobs[i], company.Name)
	}
	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)
	return company
}

func parseFarmacorpJobContent(job *Job, company string) {
	c := colly.NewCollector()
	c.OnHTML("section.aira-job-portal-company-info", func(e *colly.HTMLElement) {
		job.Content, _ = e.DOM.Html()
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s job error (%s): %v\n", company, job.Url, err)
	})
	c.Visit(job.Url)
}
