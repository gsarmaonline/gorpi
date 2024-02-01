# Go Rest APIs - Gorpi

Go Rest APIs is a thin layer built on top of the Gin framework.
The framework decides to provide few helpers and conventions which
make it very straightforward to establish simple CRUD APIs.

## Routing
Rails, the Ruby framework provides a method `resources` which defines
the RESTful routes by default for the mentioned resource.
It allows proper nesting of resources depending on where they are configured.
Though slightly laborious to get to that level of efficiency in Go, the API
layer will try to take the Rails resources as the guiding factor.

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
		svr *gorestapi.Server
		err error
	)

	if svr, err = gorestapi.New(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	if err = svr.Run(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	log.Println(svr)
}
```

## TODO
- Authenticator
- Database connection
- Statistics

