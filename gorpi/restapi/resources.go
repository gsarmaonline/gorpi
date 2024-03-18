package restapi

import (
	"fmt"

	"github.com/gauravsarma1992/go-rest-api/gorpi/routing"
)

const (
	IndexApiType  = ApiType(1)
	CreateApiType = ApiType(2)
	UpdateApiType = ApiType(3)
	ShowApiType   = ApiType(4)
	DeleteApiType = ApiType(5)
)

type (
	ResourceRoute struct {
		ResourceModel ResourceModel
		IgnoreApis    []ApiType
		ApiPrefix     string
		Version       string
	}
)

var (
	// TODO: Check if useful, otherwise remove
	DefaultApis = []ApiType{IndexApiType, CreateApiType, UpdateApiType, ShowApiType, DeleteApiType}
)

func (rRoute *ResourceRoute) GetApi() (api string) {
	ancestor := rRoute.ResourceModel.Ancestor()
	ancestorPrefix := rRoute.GetAncestorPrefix(ancestor)
	if rRoute.Version != "" {
		rRoute.ApiPrefix = fmt.Sprintf("%s/%s", rRoute.ApiPrefix, rRoute.Version)
	}
	api = fmt.Sprintf("%s%s/%s", rRoute.ApiPrefix, ancestorPrefix, rRoute.ResourceModel.String())
	return
}

func (rRoute *ResourceRoute) GetAncestorPrefix(ancestor ResourceModel) (prefix string) {
	if ancestor == nil {
		prefix = ""
		return
	}
	prefix = fmt.Sprintf("%s/%s/:%s_id",
		rRoute.GetAncestorPrefix(ancestor.Ancestor()),
		ancestor.String(),
		ancestor.String(),
	)
	return
}

func (rRoute *ResourceRoute) TranslateToRoutes(defaultHandler *DefaultHandler) (routes []*routing.Route, err error) {
	apiPath := rRoute.GetApi()
	routes = []*routing.Route{
		{
			RequestURI:    fmt.Sprintf("%s", apiPath),
			RequestMethod: "GET",
			Handler:       defaultHandler.IndexHandler,
		},
		{
			RequestURI:    fmt.Sprintf("%s/:id", apiPath),
			RequestMethod: "GET",
			Handler:       defaultHandler.ShowHandler,
		},
		{
			RequestURI:    fmt.Sprintf("%s", apiPath),
			RequestMethod: "POST",
			Handler:       defaultHandler.CreateHandler,
		},
		{
			RequestURI:    fmt.Sprintf("%s/:id", apiPath),
			RequestMethod: "PUT",
			Handler:       defaultHandler.UpdateHandler,
		},
		{
			RequestURI:    fmt.Sprintf("%s/:id", apiPath),
			RequestMethod: "DELETE",
			Handler:       defaultHandler.DeleteHandler,
		},
	}
	return
}
