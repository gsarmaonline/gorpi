package middlewares

import (
	"github.com/gin-gonic/gin"
)

func (srv *Server) RequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("example", "12345")
		c.Next()
	}
}
