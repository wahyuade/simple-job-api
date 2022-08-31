package base

import (
	"github.com/gofiber/fiber/v2"
	"wahyuade.com/simple-job-api/controllers"
	"wahyuade.com/simple-job-api/routers"
)

type Context struct {
	web         *fiber.App
	controllers controllers.Controllers
}

func (ctx Context) Web() *fiber.App {
	return ctx.web
}

func (ctx Context) Controllers() controllers.Controllers {
	return ctx.controllers
}

func (ctx Context) JWTMiddleware(c *fiber.Ctx) error {
	return nil
}

func New() Context {
	fiberApp := newFiber()
	c := controllers.NewController()
	ctx := Context{
		web:         fiberApp,
		controllers: c,
	}
	routers.Bootstrap(ctx)
	return ctx
}

func newFiber() *fiber.App {
	web := fiber.New()
	return web
}
