package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

func parseFairPlayJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "FairPlay",
		LogoUrl: "https://fairplaybo.vtexassets.com/assets/vtex/assets-builder/fairplaybo.fairplay-theme/1.0.5/imgs/logo-fairplaybo___edd19225884e46691e35a9a4df4aac8c.png",
		JobsUrl: "https://fairplay.zohorecruit.com/jobs/Grupo-Fair-Play",
		Jobs:    []Job{},
	}

	c := colly.NewCollector()
	baseUrl := "https://fairplay.zohorecruit.com/jobs/Grupo-Fair-Play/"
	c.OnHTML("input#jobs", func(e *colly.HTMLElement) {
		type Item struct {
			Name string `json:"Job_Opening_Name"`
			Id   string `json:"id"`
			City string `json:"City"`
		}
		strData := e.Attr("value")
		var items []Item
		err := json.Unmarshal([]byte(strData), &items)
		if err != nil {
			fmt.Printf("%s error: %v\n", company.Name, err)
		}
		for i := range items {
			company.Jobs = append(
				company.Jobs,
				Job{
					Title: items[i].Name,
					Url:   baseUrl + items[i].Id,
					Depto: parseDepto(items[i].City),
				},
			)
		}
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s error: %v\n", company.Name, err)
	})
	c.Visit(company.JobsUrl)

	for i := range company.Jobs {
		parseFairPlayJob(&company.Jobs[i], company.Name)
	}

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)
	return company
}

func parseFairPlayJob(job *Job, company string) {
	body, err := fetchUrl(job.Url)
	if err != nil {
		fmt.Printf("%s job error(%s): %v\n", company, job.Url, err)
	}
	line := []byte{}
	lines := bytes.Split(body, []byte("\n"))
	for i := range lines {
		if bytes.HasPrefix(lines[i], []byte("var jobs =")) {
			line = lines[i]
			break
		}
	}
	if len(line) == 0 {
		return
	}
	line = line[23:]
	line = line[:len(line)-3]

	re := regexp.MustCompile(`\\x([0-9a-fA-F]{2})`)
	jsonData := re.ReplaceAllStringFunc(string(line), func(match string) string {
		hexValue := match[2:]
		intValue, err := strconv.ParseInt(hexValue, 16, 64)
		if err != nil {
			fmt.Printf("%s job error(%s): %v\n", company, job.Url, err)
		}
		return string(rune(intValue))
	})

	type JobItem struct {
		DateOpened     string `json:"Date_Opened"`
		Industry       string `json:"Industry"`
		JobDescription string `json:"Job_Description"`
	}
	var data []JobItem
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Printf("%s job error(%s): %v\n", company, job.Url, err)
	}
	if len(data) == 0 {
		return
	}
	publishedDate, _ := time.Parse("2006-01-02", data[0].DateOpened)
	job.Content = data[0].JobDescription
	job.PublishDate = &publishedDate
	job.Area = data[0].Industry

}
