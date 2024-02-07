package routing

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func DummyHandler(c *gin.Context) {
	fmt.Println(c)
	return
}

func AddDummyRoutes(rm *RouteManager) {
	routes := []*Route{
		&Route{
			RequestURI:    "/hello",
			RequestMethod: "POST",
			Handler:       DummyHandler,
		},
		&Route{
			RequestURI:    "/hello/world",
			RequestMethod: "POST",
			Handler:       DummyHandler,
		},
		&Route{
			RequestURI:    "/hello/again",
			RequestMethod: "POST",
			Handler:       DummyHandler,
		},
		&Route{
			RequestURI:    "/hello/:id",
			RequestMethod: "GET",
			Handler:       DummyHandler,
		},
		&Route{
			RequestURI:    "/hello/:id/again/to/you",
			RequestMethod: "POST",
			Handler:       DummyHandler,
		},
	}

	for _, route := range routes {
		rm.AddRoutes(route)
	}
}

func TestRouteManagerInitialization(t *testing.T) {
	testApiEngine := gin.Default()
	rm, err := NewRouteManager(testApiEngine, nil)

	assert.NotEqual(t, rm, nil)
	assert.Equal(t, err, nil)
}

func TestRouteManagerAddRoutes(t *testing.T) {
	testApiEngine := gin.Default()
	rm, _ := NewRouteManager(testApiEngine, nil)

	AddDummyRoutes(rm)
	totalRoutes := len(rm.trie.traverse(rm.trie.rootNode, "", []string{}))

	assert.Equal(t, totalRoutes, 6)

}

func TestRouteManagerGetRoutes(t *testing.T) {
	testApiEngine := gin.Default()
	rm, _ := NewRouteManager(testApiEngine, nil)

	AddDummyRoutes(rm)

	route, err := rm.GetRoute("POST-/hello")
	assert.Equal(t, err, nil)
	assert.Equal(t, route.RequestURI, "/hello")

	route, err = rm.GetRoute("GET-/hello")
	assert.NotEqual(t, err, nil)
	fmt.Println(route)

	route, err = rm.GetRoute("GET-/hello/21")
	assert.Equal(t, err, nil)
	assert.Equal(t, route.RequestURI, "/hello/:id")
	assert.Equal(t, route.RequestMethod, "GET")

	route, err = rm.GetRoute("POST-/hello/world")
	assert.Equal(t, err, nil)
	assert.Equal(t, route.RequestURI, "/hello/world")
	assert.Equal(t, route.RequestMethod, "POST")

}
