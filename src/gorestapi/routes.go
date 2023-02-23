package gorestapi

import (
	"github.com/gin-gonic/gin"
)

type (
	Route struct {
		RequestURI         string
		RequestMethod      string
		Handler            gin.HandlerFunc
		ShouldAuthenticate bool
	}
)

func (srv *Server) setRoutes() (err error) {
	srv.AddRoute(Route{"/ping", "GET", srv.PingHandler})
	return
}

func (srv *Server) AddRoute(route Route) (err error) {
	srv.apiEngine.Handle(route.RequestMethod, route.RequestURI, route.Handler)
	return
}
