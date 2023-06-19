package gorestapi

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/gauravsarma1992/gostructs"
	"github.com/stretchr/testify/assert"
)

type TestA struct {
	A int
	B string
}

var (
	Srv *Server
)

func GetTestServer() (srv *Server) {
	if Srv == nil {
		log.Println("Refreshing server", Srv)
		Srv, _ = New(nil)
	}
	srv = Srv
	return
}

func GetTestGenerator() (g *Generator) {
	srv, _ := New(nil)
	g = srv.Generator
	g.Generate(TestA{})
	return
}

func FireCreateQuery(count int) (resps []*http.Response) {
	reqBody := []byte(`{"a": 5, "b": "hello world"}`)
	payloadB := bytes.NewBuffer(reqBody)
	for idx := 0; idx < count; idx++ {
		resp, _ := FireQuery("http://localhost:9091/api/autogen/test_a", "POST", payloadB)
		resps = append(resps, resp)
	}
	return
}

func parseHttpResponse(resp *http.Response) (respBody map[string]interface{}) {
	respBody = make(map[string]interface{})
	jsonB, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(jsonB, &respBody)
	return
}

func parseRespBody(resp *http.Response) (decodedResult *gostructs.DecodedResult) {
	respBody := parseHttpResponse(resp)
	decoder, _ := gostructs.NewDecoder(&gostructs.DecoderConfig{ShouldSnakeCase: true})
	decodedResult, _ = decoder.DecodeFreeMap(respBody["result"].(map[string]interface{}))
	return
}

func parseMultiHttpResponse(resp *http.Response) (decodedResults []*gostructs.DecodedResult) {
	respBody := parseHttpResponse(resp)
	decoder, _ := gostructs.NewDecoder(&gostructs.DecoderConfig{ShouldSnakeCase: true})
	respBodyElems := respBody["result"].([]interface{})
	for _, elem := range respBodyElems {
		decodedResult, _ := decoder.DecodeFreeMap(elem.(map[string]interface{}))
		decodedResults = append(decodedResults, decodedResult)
	}

	return
}

func FireQuery(url string, method string, payload io.Reader) (resp *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	return
}

func TestGeneratorNew(t *testing.T) {
	srv := GetTestServer()
	assert.NotNil(t, srv.Generator)
	assert.Equal(t, srv.Generator.config.ApiPrefix, DefaultGeneratorApiPrefix)
}

func TestGeneratorGenerate(t *testing.T) {
	srv := GetTestServer()
	g, _ := NewGenerator(srv, nil)
	g.Generate(TestA{3, "hello"})
	srv.Generator = g
	go g.srv.Run()

	resp, _ := FireQuery("http://localhost:9091/api/autogen/test_a", "GET", nil)
	g.srv.Cancel()

	assert.NotNil(t, g.Resources["test_a"])
	assert.Equal(t, g.Resources["test_a"].DecodedResource.Attributes["a"], 3)
	assert.Equal(t, g.Resources["test_a"].DecodedResource.Attributes["b"], "hello")
	assert.Equal(t, resp.StatusCode, 200)
}

func TestGeneratorIndex(t *testing.T) {
	g := GetTestGenerator()
	go g.srv.Run()

	FireCreateQuery(5)

	resp, _ := FireQuery("http://localhost:9091/api/autogen/test_a", "GET", nil)
	g.srv.Cancel()

	decodedResults := parseMultiHttpResponse(resp)

	assert.Equal(t, resp.StatusCode, 200)
	assert.GreaterOrEqual(t, len(decodedResults), 5)
}

func TestGeneratorCreate(t *testing.T) {
	g := GetTestGenerator()
	go g.srv.Run()
	reqBody := []byte(`{"a": 5, "b": "hello world"}`)
	payloadB := bytes.NewBuffer(reqBody)
	resp, _ := FireQuery("http://localhost:9091/api/autogen/test_a", "POST", payloadB)
	g.srv.Cancel()

	decodedResult := parseRespBody(resp)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, decodedResult.Attributes["a"], float64(5))
	assert.Equal(t, decodedResult.Attributes["b"], "hello world")
	assert.NotEqual(t, decodedResult.Attributes["id"], "")
}

func TestGeneratorShow(t *testing.T) {
	g := GetTestGenerator()
	go g.srv.Run()

	resps := FireCreateQuery(1)
	decodedResult := parseRespBody(resps[0])
	primaryKey := decodedResult.Attributes["id"].(string)

	resp, _ := FireQuery("http://localhost:9091/api/autogen/test_a/"+primaryKey, "GET", nil)
	g.srv.Cancel()

	log.Println(parseHttpResponse(resp))

	assert.Equal(t, resp.StatusCode, 200)
}
