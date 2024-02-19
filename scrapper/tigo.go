package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func parseTigoJobs() Company {
	start := time.Now()
	company := Company{
		Name:    "Tigo",
		LogoUrl: "https://images.tigocloud.net/j1bxozgharz5/boqFH1LxMcvZWBq99rTfqRe/67cff7425bdb82a76ea21f1df1ee8f7a/tigo-brand.svg",
		JobsUrl: "https://tigo.wd3.myworkdayjobs.com/wday/cxs/tigo/tigocareers/jobs",
		Jobs:    []Job{},
	}
	requestData := map[string]interface{}{
		"appliedFacets": map[string]interface{}{
			"locationCountry": []string{"db69bccc446c11de98360015c5e6daf6"},
		},
		"limit":      20,
		"offset":     0,
		"searchText": "",
	}
	type JobPost struct {
		Title         string `json:"title"`
		ExternalPath  string `json:"externalPath"`
		LocationsText string `json:"locationsText"`
	}
	type Response struct {
		JobPostings []JobPost `json:"jobPostings"`
	}

	jsonData, _ := json.Marshal(requestData)
	resp, err := http.Post(company.JobsUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("%s error: %v\n", company.Name, err)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("%s error: Non 200 response\n", company.Name)
	}
	// Response
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Printf("%s error: %v\n", company.Name, err)
	}
	jobs := make([]Job, len(response.JobPostings))
	basePath := "https://tigo.wd3.myworkdayjobs.com/wday/cxs/tigo/tigocareers"
	for i := range response.JobPostings {
		jobs[i] = Job{
			Title:   response.JobPostings[i].Title,
			DataUrl: basePath + response.JobPostings[i].ExternalPath,
			Depto:   parseDepto(response.JobPostings[i].LocationsText),
		}
		parseTigoJob(&jobs[i], company.Name)
	}
	company.Jobs = jobs

	duration := time.Since(start)
	fmt.Printf("%s: %d jobs parsed in %v\n", company.Name, len(company.Jobs), duration)
	return company
}

func parseTigoJob(job *Job, company string) {
	type JobPosting struct {
		ExternalUrl    string `json:"externalUrl"`
		JobDescription string `json:"jobDescription"`
		StartDate      string `json:"startDate"`
		Location       string `json:"location"`
	}
	type Response struct {
		JobPostingInfo JobPosting `json:"jobPostingInfo"`
	}
	body, err := fetchUrl(job.DataUrl)
	if err != nil {
		fmt.Printf("%s job error (%s): %v\n", company, job.Url, err)
	}
	resp := Response{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Printf("%s job error (%s): %v\n", company, job.Url, err)
	}
	job.Url = resp.JobPostingInfo.ExternalUrl
	job.Content = resp.JobPostingInfo.JobDescription
	job.PublishDate = parseLaPazTime("2006-01-02", resp.JobPostingInfo.StartDate)
	job.Depto = parseDepto(resp.JobPostingInfo.Location)
}
