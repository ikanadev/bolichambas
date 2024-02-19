package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func parseEconomicoJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Banco Económico",
		LogoUrl: "https://www.baneco.com.bo/images/logo.png",
		JobsUrl: "https://apiempresas.evaluatest.com/SearchEngine/job/BANCO%20ECONOMICO57/offset/0/limit/100/text/%20/language/es-MX",
		Jobs:    []Job{},
	}

	body, err := fetchUrl(company.JobsUrl)
	if err != nil {
		fmt.Printf("%s error: %v\n", company.Name, err)
	}

	data := []map[string]string{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("%s error: %v\n", company.Name, err)
	}

	jobs := make([]Job, len(data))
	baseJobUrl := "https://empresas.evaluatest.com/vacante/"
	for i := range data {
		jobs[i].Title = data[i]["name"]
		jobs[i].Depto = parseDepto(data[i]["location"])
		jobs[i].PublishDate = parseLaPazTime("2006-01-02", data[i]["publicationDate"][0:10])
		jobs[i].Url = baseJobUrl + data[i]["evaluationCode"]
		parseEconomicoJob(&jobs[i], company.Name)
	}
	company.Jobs = jobs

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)
	return company
}

func parseEconomicoJob(job *Job, company string) {
	type Data struct {
		Props struct {
			PageProps struct {
				Vacant struct {
					Vacant struct {
						Responsability            string `json:"Responsability"`
						FunctionalAreaDescription string `json:"FunctionalAreaDescription"`
					} `json:"Vacant"`
				} `json:"vacant"`
			} `json:"pageProps"`
		} `json:"props"`
	}
	c := colly.NewCollector()

	c.OnHTML("script#__NEXT_DATA__", func(e *colly.HTMLElement) {
		data := Data{}
		json.Unmarshal([]byte(e.Text), &data)
		job.Content = data.Props.PageProps.Vacant.Vacant.Responsability
		job.Area = data.Props.PageProps.Vacant.Vacant.FunctionalAreaDescription
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s job error (%s): %v\n", company, job.Url, err)
	})

	c.Visit(job.Url)
}
