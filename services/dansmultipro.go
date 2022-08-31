package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"wahyuade.com/simple-job-api/models"
)

type DansMultiPro struct {
	client *resty.Client
}

func NewDansMultiPro() DansMultiPro {
	c := resty.New()
	c = c.SetBaseURL("http://dev3.dansmultipro.co.id/api/recruitment")
	return DansMultiPro{
		client: c,
	}
}

func (d DansMultiPro) JobList(query map[string]string) (jobs []models.Job) {
	response, err := d.client.R().SetQueryParams(query).Get("/positions.json")
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(response.Body(), &jobs)
	return jobs
}

func (d DansMultiPro) JobDetail(id string) (job models.Job) {
	response, err := d.client.R().Get(fmt.Sprintf("/positions/%s", id))
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(response.Body(), &job)
	return job
}
