package main

import (
	"log"
	"os"

	"github.com/gauravsarma1992/go-rest-api/gorpi/models"
	"github.com/gauravsarma1992/go-rest-api/gorpi/restapi"
)

type (
	User struct {
		Name string `json:"name"`
	}
)

func (dm *User) String() (name string) {
	name = "users"
	return
}

func (dm *User) Ancestor() (ancestor models.ResourceModel) {
	return
}

func main() {
	var (
		mgr *restapi.RestApiManager
		err error
	)
	if mgr, err = restapi.NewRestApiManager(nil, nil); err != nil {
		log.Fatal(err)
	}
	rRoute := &restapi.ResourceRoute{
		ResourceModel: &User{},
		Version:       "v1",
	}
	mgr.AddResource(rRoute)

	if err = mgr.Run(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
