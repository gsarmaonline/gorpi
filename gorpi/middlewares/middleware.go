package middlewares

import (
	"container/list"
	"context"
	"log"

	"github.com/gauravsarma1992/go-rest-api/gorpi/api"
	"github.com/gauravsarma1992/go-rest-api/gorpi/models"
	"github.com/gin-gonic/gin"
)

type (
	Middleware interface {
		Process(*api.Request, *api.Response, *Tracker) error
	}
	MiddlewareStack struct {
		db          *models.DB
		middlewares []Middleware
		tracker     *Tracker
		ll          *list.List
	}
)

func NewMiddlewareStack(db *models.DB) (ms *MiddlewareStack) {
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

func (ms *MiddlewareStack) Exec(ctx context.Context, c *gin.Context, handler api.ApiHandlerFunc) (err error) {
	var (
		request  *api.Request
		response *api.Response
		tracker  *Tracker
	)

	request = api.NewRequest(ctx, c)
	response = api.NewResponse(request)

	if ms.db != nil {
		request.Db = ms.db
	}

	tracker = NewTracker(ms, request, response, handler)

	if err = tracker.Start(); err != nil {
		log.Println("Error in Middleware Tracker", err)
		return
	}
	return
}
