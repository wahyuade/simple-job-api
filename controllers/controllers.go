package controllers

import (
	"wahyuade.com/simple-job-api/repositories"
	"wahyuade.com/simple-job-api/services"
)

type Controllers struct {
	svc  services.Services
	repo repositories.Repositories
	Public
	Jobs
}

func NewController() Controllers {
	svc := services.NewServices()
	repo := repositories.NewRepositories()
	controller := Controllers{
		svc:  svc,
		repo: repo,
	}
	public := NewPublic(repo)
	jobs := NewJobs(svc)
	controller.Public = public
	controller.Jobs = jobs
	return controller
}
