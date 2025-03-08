package restapi

import (
	"github.com/gauravsarma1992/go-rest-api/core"
	"github.com/gauravsarma1992/go-rest-api/core/api"
	"github.com/gauravsarma1992/go-rest-api/core/routing"
	//"github.com/gin-gonic/gin"
)

type (
	BaseHandler struct {
		server *core.Server
	}
)

func NewDefaultHandler(server *core.Server) (handler *BaseHandler, err error) {
	handler = &BaseHandler{
		server: server,
	}
	return
}

func (handler *BaseHandler) IndexHandler(req *api.Request, resp *api.Response) (err error) {
	reqModel := req.Ctx.Value("route").(*routing.Route).ResourceModel
	result := []map[string]interface{}{}

	if db := req.Db.Orm.Table(reqModel.String()).Find(&result); db.Error != nil {
		resp.WriteError(db.Error)
		return
	}

	resp.Write(result)
	return
}

func (handler *BaseHandler) ShowHandler(req *api.Request, resp *api.Response) (err error) {
	reqModel := req.Ctx.Value("route").(*routing.Route).ResourceModel
	result := map[string]interface{}{}

	if db := req.Db.Orm.Table(reqModel.String()).Where("id = ?", req.Params[api.PrimaryID]).First(&result); db.Error != nil {
		resp.WriteError(db.Error)
		return
	}

	resp.Write(result)
	return
}

func (handler *BaseHandler) CreateHandler(req *api.Request, resp *api.Response) (err error) {
	return
}

func (handler *BaseHandler) UpdateHandler(req *api.Request, resp *api.Response) (err error) {
	return
}

func (handler *BaseHandler) DeleteHandler(req *api.Request, resp *api.Response) (err error) {
	return
}
