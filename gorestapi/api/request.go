package api

import (
	"github.com/gin-gonic/gin"
)

type (
	Request struct {
		GinC *gin.Context
	}
)

func NewRequest(c *gin.Context) (request *Request) {
	request = &Request{
		GinC: c,
	}
	return
}
