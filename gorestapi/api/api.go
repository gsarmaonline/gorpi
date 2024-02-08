package api

import (
	"github.com/gin-gonic/gin"
)

type (
	Request struct {
		GinC          *gin.Context
		RequestURI    string
		RequestMethod string
	}
	Response struct {
		req        *Request
		StatusCode int         `json:"status_code"`
		Body       interface{} `json:"body"`
	}
	ApiHandlerFunc func(*Request, *Response) error
)

func NewRequest(c *gin.Context) (req *Request) {
	req = &Request{
		GinC:          c,
		RequestURI:    c.Request.RequestURI,
		RequestMethod: c.Request.Method,
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
	resp.req.GinC.JSON(200, gin.H{
		"result": body,
	})
	return
}
