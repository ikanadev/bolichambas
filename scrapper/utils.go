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

func containsAnyOf(s string, list ...string) bool {
	for i := range list {
		if strings.Contains(s, list[i]) {
			return true
		}
	}
	return false
}

func parseDepto(depto string) string {
	depto = normalizeString(depto)
	if containsAnyOf(depto, "lapaz", "lpz", "alto", "payacamaya") {
		return LaPaz
	}
	if containsAnyOf(depto, "oruro") {
		return Oruro
	}
	if containsAnyOf(depto, "potos", "uyuni", "tupiza", "villazón") {
		return Potosi
	}
	if containsAnyOf(depto, "cochabamba", "cbba", "sacaba") {
		return Cochabamba
	}
	if containsAnyOf(depto, "chuquisaca", "sucre", "chuq", "padilla") {
		return Chuquisaca
	}
	if containsAnyOf(depto, "tarija", "yacuiba", "yacuíba") {
		return Tarija
	}
	if containsAnyOf(depto, "santacruz", "scz", "sanignacio", "montero") {
		return SantaCruz
	}
	if containsAnyOf(depto, "pando", "cobija") {
		return Pando
	}
	if containsAnyOf(depto, "beni", "trinidad", "be") {
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
		job.PublishDate = parseLaPazTime("2006-01-02", publishDateStr)
		job.DueDate = dueDate
	})
	c.Visit(job.Url)
}

func parseLaPazTime(layout, dateStr string) *time.Time {
	location, err := time.LoadLocation("America/La_Paz")
	if err != nil {
		return nil
	}
	date, err := time.ParseInLocation(layout, dateStr, location)
	if err != nil {
		return nil
	}
	return &date
}
