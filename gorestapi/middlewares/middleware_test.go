package middlewares

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gauravsarma1992/go-rest-api/gorestapi/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type (
	DummyMiddleware struct{}
)

func (dm *DummyMiddleware) Process(req *api.Request, resp *api.Response, tr *Tracker) (err error) {
	fmt.Println("In Dummy middleware")
	tr.Next()
	return
}

func DummyHandler(req *api.Request, resp *api.Response) (err error) {
	fmt.Println("In handler", req)
	return
}

func TestMiddlewareInit(t *testing.T) {
	ms := NewMiddlewareStack()
	ms.Add(&DummyMiddleware{})

	c := &gin.Context{
		Request: &http.Request{
			RequestURI: "/hello/world",
		},
	}
	err := ms.Exec(c, DummyHandler)
	assert.Equal(t, err, nil)
}
