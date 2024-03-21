package restapi

import (
	"log"

	"github.com/gauravsarma1992/go-rest-api/gorpi"
	"github.com/gauravsarma1992/go-rest-api/gorpi/api"
	"github.com/gauravsarma1992/go-rest-api/gorpi/routing"
	//"github.com/gin-gonic/gin"
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

	reqModel := req.Ctx.Value("route").(*routing.Route).ResourceModel
	result := []map[string]interface{}{}

	if db := req.Db.Orm.Table(reqModel.String()).Find(&result); db.Error != nil {
		log.Println("Error in fetching resource for IndexHandler", err)
		resp.WriteError(err)
	}

	resp.Write(result)
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
