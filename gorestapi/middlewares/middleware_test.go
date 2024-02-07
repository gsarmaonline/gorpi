package middlewares

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMiddlewareInit(t *testing.T) {
	ms := NewMiddlewareStack()
	ms.Add(NewLoggerMiddleware())
	c := &gin.Context{
		Request: &http.Request{
			RequestURI: "/hello/world",
		},
	}
	err := ms.Exec(c)
	assert.Equal(t, err, nil)
}
