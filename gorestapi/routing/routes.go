package routing

import (
	"errors"
	"fmt"

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
	rm.trie = NewTrie(rm.DefaultRootRoute())
	return
}

func (rm *RouteManager) DefaultRootRoute() (route *Route) {
	route = &Route{
		RequestURI:    "",
		RequestMethod: "*",
		Handler:       rm.RootHandler,
	}
	return
}

func (rm *RouteManager) RootHandler(c *gin.Context) {
	return
}

func (rm *RouteManager) AddRoutes(route *Route) (err error) {
	if _, err = rm.trie.AddPath(route); err != nil {
		return
	}
	return
}

func (rm *RouteManager) GetRoute(path string) (route *Route, err error) {
	var (
		pathNode *Node
	)
	if pathNode, err = rm.trie.GetNode(path); err != nil {
		return
	}
	if pathNode.Route == nil {
		err = errors.New("Route not found for path" + path)
		return
	}
	route = pathNode.Route
	return
}

func (route *Route) GetName() (name string) {
	name = fmt.Sprintf("%s-%s", route.RequestMethod, route.RequestURI)
	return
}
