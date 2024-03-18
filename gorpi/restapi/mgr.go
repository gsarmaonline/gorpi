package restapi

import (
	"log"

	"github.com/gauravsarma1992/go-rest-api/gorpi"
	"github.com/gauravsarma1992/go-rest-api/gorpi/routing"
)

type (
	ApiType uint8

	ResourceModel interface {
		String() string
		Ancestor() ResourceModel
	}

	RestApiManager struct {
		server *gorpi.Server

		ResourceRoutes []*ResourceRoute
		defaultHandler *DefaultHandler
	}
)

func NewRestApiManager(server *gorpi.Server) (rMgr *RestApiManager, err error) {
	if server == nil {
		if server, err = gorpi.New(nil); err != nil {
			return
		}
	}
	rMgr = &RestApiManager{
		server: server,
	}
	if rMgr.defaultHandler, err = NewDefaultHandler(server); err != nil {
		return
	}
	return
}

func (rMgr *RestApiManager) AddResource(resRoute *ResourceRoute) (err error) {
	rMgr.ResourceRoutes = append(rMgr.ResourceRoutes, resRoute)
	return
}

func (rMgr *RestApiManager) GenerateRoutes() (err error) {
	for _, rRoute := range rMgr.ResourceRoutes {
		var translatedRoutes []*routing.Route
		if translatedRoutes, err = rRoute.TranslateToRoutes(rMgr.defaultHandler); err != nil {
			return
		}
		for _, route := range translatedRoutes {
			if err = rMgr.server.RouteManager.AddRoutes(route); err != nil {
				return

			}
		}
		log.Println("Generated routes from", rRoute, " : ", translatedRoutes)
	}
	return
}

func (rMgr *RestApiManager) Run() (err error) {
	err = rMgr.server.Run()
	return
}
