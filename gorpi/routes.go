package gorpi

import (
	"github.com/gauravsarma1992/go-rest-api/gorpi/routing"
)

func (srv *Server) setRoutes() (err error) {
	if srv.RouteManager, err = routing.NewRouteManager(srv.apiEngine, srv.MiddlewareStack); err != nil {
		return
	}
	return
}
