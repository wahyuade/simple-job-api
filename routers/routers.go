package routers

import "wahyuade.com/simple-job-api/interfaces"

func Bootstrap(ctx interfaces.IContext) {
	public(ctx)
	jobs(ctx)
}
