package restapi

import (
	"github.com/gauravsarma1992/go-rest-api/gorpi"
	"github.com/gauravsarma1992/go-rest-api/gorpi/api"
)

type (
	DefaultHandler struct {
		server *gorpi.Server
	}
)

func NewDefaultHandler(server *gorpi.Server) (handler *DefaultHandler, err error) {
	handler = &DefaultHandler{
		server: server,
	}
	return
}

func (handler *DefaultHandler) IndexHandler(req *api.Request, resp *api.Response) (err error) {
	return
}

func (handler *DefaultHandler) ShowHandler(req *api.Request, resp *api.Response) (err error) {
	return
}

func (handler *DefaultHandler) CreateHandler(req *api.Request, resp *api.Response) (err error) {
	return
}

func (handler *DefaultHandler) UpdateHandler(req *api.Request, resp *api.Response) (err error) {
	return
}

func (handler *DefaultHandler) DeleteHandler(req *api.Request, resp *api.Response) (err error) {
	return
}
