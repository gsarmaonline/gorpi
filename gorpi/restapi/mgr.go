package restapi

import (
	"github.com/gauravsarma1992/go-rest-api/gorpi"
	"github.com/gauravsarma1992/go-rest-api/gorpi/routing"
)

type (
	ApiType uint8

	RestApiConfig struct {
		ApiPrefix string `json:"api_prefix"`
	}

	RestApiManager struct {
		server *gorpi.Server
		config *RestApiConfig

		ResourceRoutes []*ResourceRoute
		defaultHandler *DefaultHandler
	}
)

func NewRestApiManager(server *gorpi.Server, config *RestApiConfig) (rMgr *RestApiManager, err error) {
	if server == nil {
		if server, err = gorpi.DefaultServer(); err != nil {
			return
		}
	}
	if config == nil {
		config = DefaultRestApiConfig()
	}
	server.MiddlewareStack.Add(NewRestApiMiddleware())

	rMgr = &RestApiManager{
		server: server,
		config: config,
	}
	if rMgr.defaultHandler, err = NewDefaultHandler(server); err != nil {
		return
	}
	return
}

func DefaultRestApiConfig() (rConfig *RestApiConfig) {
	rConfig = &RestApiConfig{
		ApiPrefix: "/api",
	}
	return
}

func (rMgr *RestApiManager) AddResource(resRoute *ResourceRoute) (err error) {
	resRoute.ApiPrefix += rMgr.config.ApiPrefix
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
	}
	return
}

func (rMgr *RestApiManager) Run() (err error) {
	if err = rMgr.GenerateRoutes(); err != nil {
		return
	}
	err = rMgr.server.Run()
	return
}
