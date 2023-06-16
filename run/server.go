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
		svr *gorestapi.Server
		err error
	)

	if svr, err = gorestapi.New(nil); err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	svr.AddRoute(gorestapi.Route{"/api/success", "GET", successHandler, false})
	svr.AddRoute(gorestapi.Route{"/api/uncertain", "GET", uncertainHandler, false})
	svr.AddRoute(gorestapi.Route{"/api/failure", "GET", failureHandler, false})

	if err = svr.Run(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	log.Println(svr)
}
