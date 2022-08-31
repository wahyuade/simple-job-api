package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"wahyuade.com/simple-job-api/controllers"
)

type IContext interface {
	Web() *fiber.App
	Controllers() controllers.Controllers
	JWTMiddleware(c *fiber.Ctx) error
}
