package routers

import (
	"os"

	jwtware "github.com/gofiber/jwt/v3"
	"wahyuade.com/simple-job-api/interfaces"
)

func jobs(ctx interfaces.IContext) {
	_jobs := ctx.Web().Group("/jobs")
	_jobs.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	_jobs.Get("/", ctx.Controllers().Jobs.List)
	_jobs.Get("/:id", ctx.Controllers().Jobs.Detail)

}
