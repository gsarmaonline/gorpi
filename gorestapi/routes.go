package gorestapi

import (
	"github.com/gin-gonic/gin"
)

type (
	Authentication struct{}

	Route struct {
		RequestURI     string
		RequestMethod  string
		Handler        gin.HandlerFunc
		Authentication *Authentication

		ChildRoutes []*Route
		ParentRoute *Route
	}
)

func (route *Route) GetRequestUri() (requestUri string) {
	requestUri = route.RequestURI
	return
}

func (route *Route) GetAuthentication() (auth *Authentication) {
	auth = route.Authentication
	return
}

func (srv *Server) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (srv *Server) setRoutes() (err error) {
	srv.AddRoute(Route{
		RequestURI:     "/ping",
		RequestMethod:  "GET",
		Handler:        srv.PingHandler,
		Authentication: nil,
	})
	return
}

func (srv *Server) AddRoute(route Route) (err error) {
	srv.apiEngine.Handle(route.RequestMethod, route.GetRequestUri(), route.Handler)
	return
}
