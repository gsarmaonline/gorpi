package gorestapi

import (
	"fmt"
	"log"
	"strings"

	"github.com/gauravsarma1992/gostructs"
	"github.com/gin-gonic/gin"
)

const (
	DefaultGeneratorApiPrefix = "/api/autogen"
)

type (
	Generator struct {
		srv       *Server
		config    *GeneratorConfig
		Resources map[string]*gostructs.DecodedResult
	}
	GeneratorConfig struct {
		ApiPrefix string
	}
)

func NewGenerator(srv *Server, config *GeneratorConfig) (g *Generator, err error) {
	g = &Generator{
		srv:       srv,
		config:    config,
		Resources: make(map[string]*gostructs.DecodedResult),
	}
	if err = g.updateConfig(); err != nil {
		return
	}
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
	)
	if decodedResult, err = g.translate(resource); err != nil {
		return
	}
	g.Resources[decodedResult.Name] = decodedResult
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

func (g *Generator) getResourceAndActionFromUrl(c *gin.Context) (resourceName string, action string) {
	url := strings.Replace(c.Request.URL.Path, g.config.ApiPrefix, "", 1)
	url = strings.Trim(url, "/")
	spUrl := strings.Split(url, "/")
	resourceName = spUrl[0]
	action = c.Request.Method
	return
}

func (g *Generator) RootHandler(c *gin.Context) {
	var (
		resourceName string
		actionType   string
	)
	resourceName, actionType = g.getResourceAndActionFromUrl(c)
	log.Println(resourceName, actionType)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (g *Generator) SetupRoutes(decodedResource *gostructs.DecodedResult) (err error) {
	var (
		baseRoute string
	)
	baseRoute = fmt.Sprintf("%s/%s", g.config.ApiPrefix, decodedResource.Name)
	g.srv.AddRoute(Route{baseRoute, "GET", g.RootHandler, false})             // Index
	g.srv.AddRoute(Route{baseRoute + "/:id", "GET", g.RootHandler, false})    // Show
	g.srv.AddRoute(Route{baseRoute, "POST", g.RootHandler, false})            // Create
	g.srv.AddRoute(Route{baseRoute + "/:id", "PUT", g.RootHandler, false})    // Update
	g.srv.AddRoute(Route{baseRoute + "/:id", "DELETE", g.RootHandler, false}) // Delete
	return
}
