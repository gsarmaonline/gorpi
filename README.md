# Go Rest APIs

## Usage

To run the `go-rest-api` server

```bash
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

