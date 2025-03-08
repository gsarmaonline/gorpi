package restapi

import (
	"fmt"
	"testing"

	"github.com/gauravsarma1992/go-rest-api/core"
	"github.com/gauravsarma1992/go-rest-api/core/models"
	"github.com/stretchr/testify/assert"
)

type (
	GrandParentModel struct{}
	ParentModel      struct{}
	ChildModel       struct{}
)

func (dm *GrandParentModel) String() (name string) {
	name = "grand_parent"
	return
}

func (dm *GrandParentModel) Ancestor() (ancestor models.ResourceModel) {
	return
}

func (dm *ParentModel) String() (name string) {
	name = "parent"
	return
}

func (dm *ParentModel) Ancestor() (ancestor models.ResourceModel) {
	ancestor = &GrandParentModel{}
	return
}

func (dm *ChildModel) String() (name string) {
	name = "dummy"
	return
}

func (dm *ChildModel) Ancestor() (ancestor models.ResourceModel) {
	ancestor = &ParentModel{}
	return
}

func TestMgrInit(t *testing.T) {

	srv, err := core.DefaultServer()
	assert.Nil(t, err)

	rMgr, err := NewRestApiManager(srv, nil)
	assert.Nil(t, err)
	assert.NotNil(t, rMgr)

	return
}

func TestMgrAddRoutes(t *testing.T) {

	srv, _ := core.DefaultServer()

	rMgr, _ := NewRestApiManager(srv, nil)

	rRoute := &ResourceRoute{
		ResourceModel: &ChildModel{},
		Version:       "v1",
	}
	rMgr.AddResource(rRoute)
	fmt.Println(rMgr.GenerateRoutes())

	return
}
