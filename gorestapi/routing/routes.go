package routing

import (
	"github.com/gin-gonic/gin"
)

type (
	Authentication struct{}

	RouteManager struct {
		apiEngine *gin.Engine
		trie      *Trie
	}

	Route struct {
		RequestURI     string
		RequestMethod  string
		Handler        gin.HandlerFunc
		Authentication *Authentication

		ChildRoutes []*Route
		ParentRoute *Route
	}
)

func NewRouteManager(apiEngine *gin.Engine) (rm *RouteManager, err error) {
	rm = &RouteManager{
		apiEngine: apiEngine,
	}
	rm.trie = NewTrie(rm.RootHandler())
	return
}

func (rm *RouteManager) DefaultRootRoute() (route *Route) {
	route = &Route{
		RequestURI:    "/",
		RequestMethod: "*",
		Handler:       rm.RootHandler,
	}
	return
}

func (rm *RouteManager) RootHandler(c *gin.Context) {
	return
}

func (rm *RouteManager) AddRoutes(route *Route) (err error) {
	rm.findLongestPrefix(route)
	return
}

func (rm *RouteManager) findLongestPrefix(route *Route) (nearestRoute *Route, err error) {
	return
}
