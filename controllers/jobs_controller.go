package controllers

import (
	"github.com/gofiber/fiber/v2"
	"wahyuade.com/simple-job-api/helpers"
	"wahyuade.com/simple-job-api/models"
	"wahyuade.com/simple-job-api/services"
)

type Jobs struct {
	svc services.Services
}

func (j Jobs) List(c *fiber.Ctx) error {
	query := helpers.GetMappedQuery(c.Context().QueryArgs().String())
	list_rule := make(map[string][]string)
	list_rule["description"] = []string{"string"}
	list_rule["location"] = []string{"string"}
	list_rule["full_time"] = []string{"boolean"}
	list_rule["page"] = []string{"number"}
	status, errValidation, cleanQuery := helpers.Validate(query, list_rule)
	if !status {
		return c.Status(400).JSON(models.APIResponse{
			Status:  false,
			Message: "Invalid request",
			Error:   errValidation,
		})
	}

	cleanQueryAsResty := make(map[string]string)
	for k, v := range cleanQuery {
		if v == nil {
			continue
		}
		cleanQueryAsResty[k] = v.(string)
	}

	return c.JSON(models.APIResponse{
		Status:  true,
		Message: "Success",
		Data:    j.svc.DansMultiPro.JobList(cleanQueryAsResty),
	})
}

func (j Jobs) Detail(c *fiber.Ctx) error {
	id := c.Params("id", "")
	job := j.svc.DansMultiPro.JobDetail(id)
	if job.Id == "" {
		return c.JSON(models.APIResponse{
			Status:  false,
			Message: "Not found",
		})
	}
	return c.JSON(models.APIResponse{
		Status:  true,
		Message: "Success",
		Data:    job,
	})
}

func NewJobs(svc services.Services) Jobs {
	return Jobs{
		svc: svc,
	}
}
