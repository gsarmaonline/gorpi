# Go Rest APIs - Gorpi

Gorpi is a thin layer built on top of the Gin framework.
The framework decides to provide few helpers and conventions which
make it very straightforward to establish simple CRUD APIs.

## Routing
Rails, the Ruby framework provides a method `resources` which defines
the RESTful routes by default for the mentioned resource.
It allows proper nesting of resources depending on where they are configured.
Though slightly laborious to get to that level of efficiency in Go, the API
layer will try to take the Rails resources as the guiding factor.
```
	Route struct {
		RequestURI     string
		RequestMethod  string
		Handler        gin.HandlerFunc
		Authentication *Authentication

		ChildRoutes []*Route
		ParentRoute *Route
	}
```

## Database connections
Gorpi provides a database library on top of gorm to provide easy access
to an ORM via the server object.

## Usage

To run the `go-rest-api` server

```golang
package main

import (
	"log"
	"os"

	gorestapi "github.com/gauravsarma1992/gorestapi"
)

func main() {
	var (
		srv *gorestapi.Server
		err error
	)

	if srv, err = gorestapi.New(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	if err = srv.Run(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	log.Println(srv)
}
```

## TODO
- Authenticator
- Database connection
- Statistics

