package main

import (
	"log"
	"os"

	gorestapi "github.com/gauravsarma1992/go-rest-api/gorestapi"
	"github.com/gauravsarma1992/go-rest-api/gorestapi/routing"
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

func addDummyRoutes(rm *routing.RouteManager) {

	routes := []*routing.Route{
		&routing.Route{
			RequestURI:    "/hello",
			RequestMethod: "POST",
			Handler:       successHandler,
		},
		&routing.Route{
			RequestURI:    "/hello/world",
			RequestMethod: "POST",
			Handler:       successHandler,
		},
		&routing.Route{
			RequestURI:    "/hello/again",
			RequestMethod: "POST",
			Handler:       successHandler,
		},
		&routing.Route{
			RequestURI:    "/hello/:id",
			RequestMethod: "GET",
			Handler:       successHandler,
		},
		&routing.Route{
			RequestURI:    "/hello/:id/again/to/you",
			RequestMethod: "POST",
			Handler:       successHandler,
		},
	}

	for _, route := range routes {
		rm.AddRoutes(route)
	}
}

func main() {
	var (
		srv *gorestapi.Server
		cfg *gorestapi.Config
		err error
	)

	cfg = &gorestapi.Config{}

	cfg.Server.Host = "127.0.0.1"
	cfg.Server.Port = "8090"
	cfg.Database.Username = "root"
	cfg.Database.Password = ""
	cfg.Database.DbName = "gorpi"
	cfg.Database.Host = "127.0.0.1"
	cfg.Database.Port = "3306"

	if srv, err = gorestapi.New(cfg); err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	addDummyRoutes(srv.RouteManager)

	if err = srv.Run(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	log.Println(srv)
}
