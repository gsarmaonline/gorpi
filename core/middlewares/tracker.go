package middlewares

import (
	"container/list"

	"github.com/gauravsarma1992/go-rest-api/core/api"
)

type (
	Tracker struct {
		ms *MiddlewareStack

		req     *api.Request
		resp    *api.Response
		handler api.ApiHandlerFunc

		currNode *list.Element
	}
)

func NewTracker(ms *MiddlewareStack, req *api.Request, resp *api.Response, handler api.ApiHandlerFunc) (tr *Tracker) {
	tr = &Tracker{
		ms:      ms,
		req:     req,
		resp:    resp,
		handler: handler,
	}
	return

}

func (tr *Tracker) Start() (err error) {
	tr.currNode = tr.ms.ll.Front()
	err = tr.Exec()
	return
}

func (tr *Tracker) Next() (err error) {
	// Having the next node as null means that
	// we have reached the end of the list and
	// we can call the actual route handler
	if tr.currNode.Next() == nil {
		if err = tr.handler(tr.req, tr.resp); err != nil {
			return
		}
		return
	}
	tr.currNode = tr.currNode.Next()
	if err = tr.Exec(); err != nil {
		return
	}
	return
}

func (tr *Tracker) Exec() (err error) {
	if err = tr.currNode.Value.(Middleware).Process(tr.req, tr.resp, tr); err != nil {
		return
	}
	return
}
