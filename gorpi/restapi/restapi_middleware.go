package restapi

import (
	"log"

	"github.com/gauravsarma1992/go-rest-api/gorpi/api"
	"github.com/gauravsarma1992/go-rest-api/gorpi/middlewares"
)

type (
	RestApiMiddleware struct {
	}
)

func NewRestApiMiddleware() (rmM *RestApiMiddleware) {
	rmM = &RestApiMiddleware{}
	return
}

func (rmM *RestApiMiddleware) Process(req *api.Request, resp *api.Response, tr *middlewares.Tracker) (err error) {
	log.Println(req.Ctx.Value("route"))
	tr.Next()
	return
}
