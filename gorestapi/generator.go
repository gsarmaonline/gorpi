package gorestapi

import (
	"github.com/gauravsarma1992/gostructs"
)

type (
	Generator struct {
		Resources map[string]map[string]interface{}
	}
)

func NewGenerator() (g *Generator, err error) {
	g = &Generator{
		Resources: make(map[string]map[string]interface{}),
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
	g.Resources[decodedResult.Name] = decodedResult.Attributes
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

func (g *Generator) SetupRoutes() (err error) {
	return
}
