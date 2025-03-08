package routing

import (
	"context"
	"fmt"

	"github.com/gauravsarma1992/go-rest-api/core/api"
	"github.com/gauravsarma1992/go-rest-api/core/middlewares"
	"github.com/gauravsarma1992/go-rest-api/core/models"
	"github.com/gin-gonic/gin"
)

const (
	ContextRouteKey = "route"
)

type (
	RouteManager struct {
		apiEngine      *gin.Engine
		middlwareStack *middlewares.MiddlewareStack
	}

	Route struct {
		RequestURI    string
		RequestMethod string
		Handler       api.ApiHandlerFunc
		ResourceModel models.ResourceModel
		Params        map[string]string
	}
)

func NewRouteManager(apiEngine *gin.Engine, ms *middlewares.MiddlewareStack) (rm *RouteManager, err error) {
	rm = &RouteManager{
		apiEngine:      apiEngine,
		middlwareStack: ms,
	}
	// The noroute handler handles the routes for routes which are
	// not defined. Since we are not defining any routes on the
	// gin context, everything will be handled by the root handler
	rm.apiEngine.NoRoute(rm.RootHandler)
	return
}

func (rm *RouteManager) GetDefaultBaseHandler() (route *Route) {
	route = &Route{
		RequestURI:    "/",
		RequestMethod: "*",
		Handler:       rm.BaseHandler,
	}
	return
}

func (rm *RouteManager) BaseHandler(req *api.Request, resp *api.Response) (err error) {
	return
}

func (rm *RouteManager) RootHandler(c *gin.Context) {
	var (
		route *Route
		err   error
		ctx   context.Context
	)

	path := fmt.Sprintf("%s-%s", c.Request.Method, c.Request.RequestURI)
	if route, err = rm.GetRoute(path); err != nil {
		c.JSON(400, gin.H{
			"request": path,
			"message": "Route not found - " + err.Error(),
		})
		return
	}

	ctx = context.WithValue(context.Background(), ContextRouteKey, route)

	if err = rm.middlwareStack.Exec(ctx, c, route.Handler, route.Params); err != nil {
		c.JSON(500, gin.H{
			"request": path,
			"message": "Handler errored - " + err.Error(),
		})
	}
	return
}

func (rm *RouteManager) AddRoutes(route *Route) (err error) {
	return
}

func (rm *RouteManager) GetRoute(path string) (route *Route, err error) {
	return
}

func (route *Route) GetName() (name string) {
	name = fmt.Sprintf("%s-%s", route.RequestMethod, route.RequestURI)
	return
}
