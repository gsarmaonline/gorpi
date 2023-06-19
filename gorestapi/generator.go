package gorestapi

import (
	"context"
	"errors"
	"fmt"
	"strings"

	godblMongo "github.com/gauravsarma1992/godbl/godbl/adapters/mongo"
	godblResource "github.com/gauravsarma1992/godbl/godbl/resource"
	"github.com/gauravsarma1992/gostructs"
	"github.com/gin-gonic/gin"
)

const (
	DefaultGeneratorApiPrefix = "/api/autogen"
)

type (
	Generator struct {
		srv       *Server
		db        godblResource.Db
		config    *GeneratorConfig
		Resources map[string]*ResourceInfo
	}
	ResourceInfo struct {
		Orig            interface{}
		DecodedResource *gostructs.DecodedResult
	}
	GeneratorConfig struct {
		ApiPrefix string
	}
)

func NewGenerator(srv *Server, config *GeneratorConfig) (g *Generator, err error) {
	g = &Generator{
		srv:       srv,
		config:    config,
		Resources: make(map[string]*ResourceInfo),
	}
	if err = g.updateConfig(); err != nil {
		return
	}
	if err = g.updateDbConfig(); err != nil {
		return
	}
	return
}

func (g *Generator) updateDbConfig() (err error) {
	var (
		mongodb *godblMongo.MongoDb
	)
	if mongodb, err = godblMongo.NewMongoDb(context.TODO(), nil); err != nil {
		return
	}
	g.db = godblResource.Db(mongodb)
	return
}

func (g *Generator) updateConfig() (err error) {
	if g.config == nil {
		g.config = &GeneratorConfig{}
	}
	if g.config.ApiPrefix == "" {
		g.config.ApiPrefix = DefaultGeneratorApiPrefix
	}
	return
}

func (g *Generator) Generate(resource interface{}) (err error) {
	var (
		decodedResult *gostructs.DecodedResult
		resourceInfo  *ResourceInfo
	)
	if decodedResult, err = g.translate(resource); err != nil {
		return
	}
	resourceInfo = &ResourceInfo{
		Orig:            resource,
		DecodedResource: decodedResult,
	}
	g.Resources[decodedResult.Name] = resourceInfo

	if err = g.SetupRoutes(decodedResult); err != nil {
		return
	}
	return
}

func (g *Generator) translate(resource interface{}) (result *gostructs.DecodedResult, err error) {
	var (
		decoder *gostructs.Decoder
	)
	decoder, _ = gostructs.NewDecoder(&gostructs.DecoderConfig{ShouldSnakeCase: true})
	if result, err = decoder.Decode(resource); err != nil {
		return
	}
	return
}

func (g *Generator) getResourceName(c *gin.Context) (resourceName string) {
	url := strings.Replace(c.Request.URL.Path, g.config.ApiPrefix, "", 1)
	url = strings.Trim(url, "/")
	spUrl := strings.Split(url, "/")
	resourceName = spUrl[0]
	return
}

func (g *Generator) GetResource(c *gin.Context) (resourceInfo *ResourceInfo, err error) {
	resourceName := g.getResourceName(c)
	resourceInfo = &ResourceInfo{}
	isPresent := true

	if resourceInfo, isPresent = g.Resources[resourceName]; isPresent == false {
		err = errors.New("Resource not found")
		return
	}
	return
}

func (g *Generator) RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (g *Generator) IndexHandler(c *gin.Context) {
	var (
		err      error
		resource *ResourceInfo
		result   []godblResource.Resource
	)
	if resource, err = g.GetResource(c); err != nil {
		ResourceNotFoundHandler(c, "")
		return
	}
	if result, err = g.db.FindMany(resource.DecodedResource); err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"result":  result,
	})
}

func (g *Generator) CreateHandler(c *gin.Context) {
	var (
		err      error
		resource *ResourceInfo
		result   godblResource.Resource
	)
	if resource, err = g.GetResource(c); err != nil {
		ResourceNotFoundHandler(c, "")
		return
	}
	if err = c.ShouldBindJSON(&resource.DecodedResource.Attributes); err != nil {
		RequestBodyClientErrorHandler(c, err)
		return
	}
	if result, err = g.db.InsertOne(resource.DecodedResource); err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"result":  result,
	})
}

func (g *Generator) ShowHandler(c *gin.Context) {
	var (
		err      error
		resource *ResourceInfo
		result   []godblResource.Resource
	)
	if resource, err = g.GetResource(c); err != nil {
		ResourceNotFoundHandler(c, "")
		return
	}
	resourceId := c.Param("id")
	resource.DecodedResource.Attributes["_id"] = resourceId
	if result, err = g.db.FindMany(resource.DecodedResource); err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"result":  result,
	})
}

func (g *Generator) UpdateHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (g *Generator) DeleteHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (g *Generator) SetupRoutes(decodedResource *gostructs.DecodedResult) (err error) {
	var (
		baseRoute string
	)
	baseRoute = fmt.Sprintf("%s/%s", g.config.ApiPrefix, decodedResource.Name)
	g.srv.AddRoute(Route{baseRoute, "GET", g.IndexHandler, false})              // Index
	g.srv.AddRoute(Route{baseRoute + "/:id", "GET", g.ShowHandler, false})      // Show
	g.srv.AddRoute(Route{baseRoute, "POST", g.CreateHandler, false})            // Create
	g.srv.AddRoute(Route{baseRoute + "/:id", "PUT", g.UpdateHandler, false})    // Update
	g.srv.AddRoute(Route{baseRoute + "/:id", "DELETE", g.DeleteHandler, false}) // Delete
	return
}
