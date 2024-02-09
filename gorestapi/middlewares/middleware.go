package middlewares

import (
	"container/list"
	"fmt"
	"github.com/gauravsarma1992/go-rest-api/gorestapi/api"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Middleware interface {
		Process(*api.Request, *api.Response, *Tracker) error
	}
	MiddlewareStack struct {
		db          *gorm.DB
		middlewares []Middleware
		tracker     *Tracker
		ll          *list.List
	}
)

func NewMiddlewareStack(db *gorm.DB) (ms *MiddlewareStack) {
	ms = &MiddlewareStack{
		ll: list.New(),
		db: db,
	}
	ms.Add(NewLoggerMiddleware())
	return
}

func (ms *MiddlewareStack) Add(middleware Middleware) {
	ms.middlewares = append(ms.middlewares, middleware)
	ms.ll.PushBack(middleware)
	return
}

func (ms *MiddlewareStack) Exec(c *gin.Context, handler api.ApiHandlerFunc) (err error) {
	var (
		request  *api.Request
		response *api.Response
		tracker  *Tracker
	)

	request = api.NewRequest(c)
	response = api.NewResponse(request)

	if ms.db != nil {
		request.Db = ms.db
	}

	tracker = NewTracker(ms, request, response, handler)

	if err = tracker.Start(); err != nil {
		fmt.Println(err)
		return
	}
	return
}
