package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

const (
	LaPaz      = "La Paz"
	Oruro      = "Oruro"
	Potosi     = "Potosí"
	Cochabamba = "Cochabamba"
	Chuquisaca = "Chuquisaca"
	Tarija     = "Tarija"
	SantaCruz  = "Santa Cruz"
	Pando      = "Pando"
	Beni       = "Beni"
)

func normalizeString(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ToLower(s)
	return s
}

func parseDepto(depto string) string {
	depto = normalizeString(depto)
	if strings.Contains(depto, "lapaz") {
		return LaPaz
	}
	if strings.Contains(depto, "oruro") {
		return Oruro
	}
	if strings.Contains(depto, "potos") {
		return Potosi
	}
	if strings.Contains(depto, "uyuni") {
		return Potosi
	}
	if strings.Contains(depto, "cochabamba") {
		return Cochabamba
	}
	if strings.Contains(depto, "chuquisaca") {
		return Chuquisaca
	}
	if strings.Contains(depto, "sucre") {
		return Chuquisaca
	}
	if strings.Contains(depto, "tarija") {
		return Tarija
	}
	if strings.Contains(depto, "santacruz") {
		return SantaCruz
	}
	if strings.Contains(depto, "pando") {
		return Pando
	}
	if strings.Contains(depto, "beni") {
		return Beni
	}
	if strings.Contains(depto, "trinidad") {
		return Beni
	}
	return ""
}

func fetchUrl(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// The company needs to have a url and a name
func parseEvaluar(company *Company) {
	type Data struct {
		Html string `json:"html"`
	}
	// fetch the jobs
	body, err := fetchUrl(company.JobsUrl)
	if err != nil {
		fmt.Printf("%s Error: %v\n", company.Name, err)
	}
	// parse the jobs
	data := Data{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("%s Error: %v\n", company.Name, err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data.Html))
	if err != nil {
		fmt.Printf("%s Error: %v\n", company.Name, err)
	}
	doc.Find("a.job_listing").Each(func(_ int, s *goquery.Selection) {
		title := strings.ReplaceAll(s.Find("h4").Text(), "\n", "")
		title = strings.TrimSpace(title)
		depto := s.Find("ul").Find("li:nth-child(2)").Text()
		company.Jobs = append(company.Jobs, Job{
			Title: title,
			Url:   s.AttrOr("href", ""),
			Depto: parseDepto(depto),
		})
	})

	for i := range company.Jobs {
		parseEvaluarDetails(&company.Jobs[i], company.Name)
	}
}

func parseEvaluarDetails(job *Job, company string) {
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
