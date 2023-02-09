package main

import (
	"log"
	"os"

	ginboilerplate "github.com/gauravsarma1992/ginboilerplate"
)

func main() {
	var (
		svr *ginboilerplate.Server
		err error
	)

	if svr, err = ginboilerplate.New(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	if err = svr.Run(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	log.Println(svr)
}
