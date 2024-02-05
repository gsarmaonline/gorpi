package gorestapi

import (
	"github.com/gauravsarma1992/go-rest-api/gorestapi/routing"
)

func (srv *Server) setRoutes() (err error) {
	if srv.RouteManager, err = routing.NewRouteManager(srv.apiEngine); err != nil {
		return
	}
	return
}
