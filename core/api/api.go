package api

import (
	"context"
	"log"

	"github.com/gauravsarma1992/go-rest-api/core/models"
	"github.com/gin-gonic/gin"
)

const (
	Debug     = true
	PrimaryID = ":id"
)

type (
	Request struct {
		Ctx           context.Context
		GinC          *gin.Context
		RequestURI    string
		RequestMethod string

		Params map[string]string

		Db *models.DB
	}
	Response struct {
		req        *Request
		StatusCode int         `json:"status_code"`
		Body       interface{} `json:"body"`
	}
	ApiHandlerFunc func(*Request, *Response) error
)

func NewRequest(ctx context.Context, c *gin.Context, params map[string]string) (req *Request) {
	req = &Request{
		Ctx:           ctx,
		GinC:          c,
		RequestURI:    c.Request.RequestURI,
		RequestMethod: c.Request.Method,
		Params:        params,
	}
	return
}

func NewResponse(req *Request) (resp *Response) {
	resp = &Response{
		req: req,
	}
	return
}

func (resp *Response) Write(body interface{}) {
	resp.WriteJSON(200, body)
	return
}

func (resp *Response) WriteJSON(statusCode int, body interface{}) {
	resp.req.GinC.JSON(statusCode, gin.H{
		"result": body,
	})

	return
}

func (resp *Response) WriteError(err error) {
	if Debug {
		log.Println("Error in request: ", resp.req.RequestURI, ". Failed with error ->", err)
	}
	resp.WriteJSON(500, err.Error())
	return
}
