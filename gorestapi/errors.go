package gorestapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ResourceNotFoundHandler(c *gin.Context, resourceName string) {
	c.JSON(404, gin.H{"error": fmt.Sprintf("Resource %s not found", resourceName)})
}

func RequestBodyClientErrorHandler(c *gin.Context, err error) {
	c.JSON(400, gin.H{"error": fmt.Sprintf("Request body client error: %s", err)})
}
