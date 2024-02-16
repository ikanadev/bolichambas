package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func parseDhlJobs() Company {
	/* Original JS function to parse title
	var getFormattedTitle = function(e){
		var t;
		if(e){
			var t=e.replace(/[\$_|`$\-+:,\/#&\[\]@\{\}*%.()?â€“']/g,"-").replace(/ /g,"-").replace(/-+/g,"-");
			"-"==t.charAt(t.length-1)&&(t=t.substring(0,t.length-1))
		}
		return t||""
	}
	*/
	// Go format title JS equivalent
	getFormattedTitle := func(e string) string {
		if e == "" {
			return ""
		}
		var t string
		re := regexp.MustCompile(`[\$_|` + "`" + `\-+:,\/#&\[\]@\{\}*%.()?â€“']+`)
		t = re.ReplaceAllString(e, "-")
		t = strings.ReplaceAll(t, " ", "-")
		t = regexp.MustCompile(`-+`).ReplaceAllString(t, "-")

		if len(t) > 0 && t[len(t)-1] == '-' {
			t = t[:len(t)-1]
		}
		return t
	}
	const reqData = `{
		"lang": "es_amer",
		"deviceType": "desktop",
		"country": "amer",
		"pageName": "search-results",
		"ddoKey": "eagerLoadRefineSearch",
		"sortBy": "",
		"subsearch": "",
		"from": 0,
		"jobs": true,
		"counts": true,
		"all_fields": [
			"country",
			"phLocSlider",
			"category",
			"careerLevel",
			"workHours",
			"businessUnit",
			"contractType"
		],
		"size": 10,
		"clearAll": false,
		"jdsource": "facets",
		"isSliderEnable": true,
		"pageId": "page17",
		"siteType": "external",
		"keywords": "",
		"global": true,
		"selected_fields": {
			"country": [
				"Bolivia"
			]
		},
		"locationData": {
			"sliderRadius": 105,
			"aboveMaxRadius": true,
			"LocationUnit": "kilometers"
		},
		"s": "1"
	}`
	start := time.Now()
	company := Company{
		Name:    "DHL",
		LogoUrl: "https://cdn.phenompeople.com/CareerConnectResources/DPDHGLOBAL/es_amer/desktop/assets/images/v-1672047183342-header_logo.png",
		JobsUrl: "https://careers.dhl.com/widgets",
		Jobs:    []Job{},
	}

	resp, err := http.Post(company.JobsUrl, "application/json", bytes.NewBuffer([]byte(reqData)))
	if err != nil {
		fmt.Printf("%s error: %v\n", company.Name, err)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("%s error: Non 200 response\n", company.Name)
	}

	type JobPost struct {
		Title      string `json:"title"`
		JobSeqNo   string `json:"jobSeqNo"`
		PostedDate string `json:"postedDate"`
		Category   string `json:"category"`
	}
	baseUrl := "https://careers.dhl.com/amer/es/job/"
	var data struct {
		EagerLoadRefineSearch struct {
			Data struct {
				Jobs []JobPost `json:"jobs"`
			} `json:"data"`
		} `json:"eagerLoadRefineSearch"`
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Printf("%s error: %v\n", company.Name, err)
	}

	jobs := make([]Job, len(data.EagerLoadRefineSearch.Data.Jobs))
	for i, respJob := range data.EagerLoadRefineSearch.Data.Jobs {
		publishDate, _ := time.Parse("2006-01-02", respJob.PostedDate[:10])
		jobs[i] = Job{
			Title:       respJob.Title,
			Area:        respJob.Category,
			Url:         baseUrl + respJob.JobSeqNo + "/" + getFormattedTitle(respJob.Title),
			PublishDate: &publishDate,
		}
	}
	company.Jobs = jobs

	for i := range company.Jobs {
		parseDhlJob(&company.Jobs[i], company.Name)
	}

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)
	return company
}
func parseDhlJob(job *Job, company string) {
	c := colly.NewCollector()

	found := false
	c.OnHTML("script[type='application/ld+json']", func(e *colly.HTMLElement) {
		if found {
			return
		}
		fmt.Println("Script found")
		found = true
		var data struct {
			Description string `json:"description"`
		}
		json.Unmarshal([]byte(e.Text), &data)
		job.Content = data.Description
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("%s job error (%s): %v\n", company, job.Url, err)
	})
	c.Visit(job.Url)
}
