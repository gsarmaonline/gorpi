package api

import "github.com/gin-gonic/gin"

type (
	Request struct {
		GinC          *gin.Context
		RequestURI    string
		RequestMethod string
	}
	Response struct {
		req *Request
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

func (resp *Response) Write(c *gin.Context) {
	return
}
