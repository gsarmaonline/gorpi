package middlewares

import (
	"github.com/gauravsarma1992/go-rest-api/gorestapi/api"
	"gorm.io/gorm"
)

type (
	Middleware interface {
		Process(*api.Request, *api.Response)
	}
	MiddlewareStack struct {
		db          *gorm.DB
		middlewares []Middleware
	}
)

func NewMiddlewareStack() (ms *MiddlewareStack) {
	ms = &MiddlewareStack{}
	return
}

func (ms *MiddlewareStack) Add(middleware Middleware) {
	ms.middlewares = append(ms.middlewares, middleware)
	return
}

func (ms *MiddlewareStack) Exec() {
	request := &api.Request{}
	response := &api.Response{}

	for _, middleware := range ms.middlewares {
		middleware.Process(request, response)
	}
}
