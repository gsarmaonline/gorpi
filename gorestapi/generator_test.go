package gorestapi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetTestServer() (srv *Server) {
	srv, _ = New(nil)
	return
}

func GetTestGenerator() (g *Generator) {
	type TestA struct {
		A int
		B string
	}
	srv, _ := New(nil)
	g = srv.Generator
	g.Generate(TestA{A: 3, B: "hello"})
	return
}

func FireQuery(url string, method string, payload interface{}) (resp *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)
	resp, err = client.Do(req)
	return
}

func TestGeneratorNew(t *testing.T) {
	srv := GetTestServer()
	assert.NotNil(t, srv.Generator)
	assert.Equal(t, srv.Generator.config.ApiPrefix, DefaultGeneratorApiPrefix)
}

func TestGeneratorGenerate(t *testing.T) {
	g := GetTestGenerator()
	go g.srv.Run()

	resp, _ := FireQuery("http://localhost:9091/api/autogen/test_a", "GET", nil)
	g.srv.Cancel()

	assert.NotNil(t, g.Resources["test_a"])
	assert.Equal(t, g.Resources["test_a"].Attributes["a"], 3)
	assert.Equal(t, g.Resources["test_a"].Attributes["b"], "hello")
	assert.Equal(t, resp.StatusCode, 200)
}

func TestGeneratorIndex(t *testing.T) {
	g := GetTestGenerator()
	go g.srv.Run()

	resp, _ := FireQuery("http://localhost:9091/api/autogen/test_a", "GET", nil)
	g.srv.Cancel()

	assert.Equal(t, resp.StatusCode, 200)
}

func TestGeneratorShow(t *testing.T) {
	g := GetTestGenerator()
	go g.srv.Run()

	resp, _ := FireQuery("http://localhost:9091/api/autogen/test_a/1", "GET", nil)
	g.srv.Cancel()

	assert.Equal(t, resp.StatusCode, 200)
}
