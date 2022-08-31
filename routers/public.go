package routers

import "wahyuade.com/simple-job-api/interfaces"

func public(ctx interfaces.IContext) {
	ctx.Web().Post("/register", ctx.Controllers().Register)
	ctx.Web().Post("/login", ctx.Controllers().Login)
}
