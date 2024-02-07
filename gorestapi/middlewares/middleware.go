package middlewares

import (
	"container/list"
	"github.com/gauravsarma1992/go-rest-api/gorestapi/api"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Middleware interface {
		Process(*api.Request, *list.Element) (*api.Response, error)
	}
	MiddlewareStack struct {
		db          *gorm.DB
		middlewares []Middleware
		tracker     *list.List
	}
)

func NewMiddlewareStack() (ms *MiddlewareStack) {
	ms = &MiddlewareStack{
		tracker: list.New(),
	}
	ms.Add(NewLoggerMiddleware())
	return
}

func (ms *MiddlewareStack) Add(middleware Middleware) {
	ms.middlewares = append(ms.middlewares, middleware)
	ms.tracker.PushBack(middleware)
	return
}

func (ms *MiddlewareStack) Exec(c *gin.Context) (err error) {
	var (
		request  *api.Request
		response *api.Response
	)

	request = api.NewRequest(c)
	response = api.NewResponse(request)

	startElem := ms.tracker.Front()

	startElem.Value.(Middleware).Process(request, startElem)

	response.Write(c)
	return
}
