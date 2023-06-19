package main

import (
	"log"
	"os"

	gorestapi "github.com/gauravsarma1992/go-rest-api/gorestapi"
	"github.com/gin-gonic/gin"
)

var (
	count int
)

func uncertainHandler(c *gin.Context) {
	statusCode := 500
	message := "failure"
	if count%2 == 0 {
		statusCode = 200
		message = "success"
	}
	count += 1
	c.JSON(statusCode, gin.H{
		"message": message,
	})
	return
}

func failureHandler(c *gin.Context) {
	c.JSON(500, gin.H{
		"message": "failed",
	})
	return
}

func successHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

func main() {
	var (
		srv *gorestapi.Server
		err error
	)

	if srv, err = gorestapi.New(nil); err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	srv.AddRoute(gorestapi.Route{"/api/success", "GET", successHandler, false})
	srv.AddRoute(gorestapi.Route{"/api/uncertain", "GET", uncertainHandler, false})
	srv.AddRoute(gorestapi.Route{"/api/failure", "GET", failureHandler, false})

	if err = srv.Run(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	log.Println(srv)
}
