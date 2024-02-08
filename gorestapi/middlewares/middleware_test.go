package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gauravsarma1992/go-rest-api/gorestapi/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type (
	DummyMiddleware        struct{}
	InterruptingMiddleware struct{}
)

func (dm *DummyMiddleware) Process(req *api.Request, resp *api.Response, tr *Tracker) (err error) {
	fmt.Println("In Dummy middleware")
	tr.Next()
	return
}

func (dm *InterruptingMiddleware) Process(req *api.Request, resp *api.Response, tr *Tracker) (err error) {
	resp.Write("Interrupting in middleware")
	return
}

func DummyHandler(req *api.Request, resp *api.Response) (err error) {
	fmt.Println("In handler", req)
	resp.Write(req.RequestURI)
	return
}

func DummyContext() (c *gin.Context) {
	w := httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request = &http.Request{
		RequestURI: "/hello/world",
		Method:     "POST",
	}
	return
}

func TestMiddlewareInit(t *testing.T) {
	ms := NewMiddlewareStack()
	ms.Add(&DummyMiddleware{})

	c := DummyContext()

	err := ms.Exec(c, DummyHandler)
	assert.Equal(t, err, nil)
}

func TestMiddlewareInterruption(t *testing.T) {
	ms := NewMiddlewareStack()
	ms.Add(&InterruptingMiddleware{})

	c := DummyContext()

	err := ms.Exec(c, DummyHandler)
	assert.Equal(t, err, nil)
}
