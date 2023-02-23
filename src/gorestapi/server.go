package gorestapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	DefaultConfigFile = "./config.json"
)

type (
	Server struct {
		apiEngine  *gin.Engine
		ConfigFile string
		Config     *Config
	}

	Config struct {
		Server struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"server"`
	}
)

func New() (srv *Server, err error) {
	srv = &Server{
		apiEngine: gin.Default(),
		Config:    &Config{},
	}
	if err = srv.setConfig(); err != nil {
		return
	}
	if err = srv.setRoutes(); err != nil {
		return
	}
	return
}

func (srv *Server) setConfig() (err error) {
	var (
		contB []byte
	)
	srv.ConfigFile = DefaultConfigFile
	if contB, err = ioutil.ReadFile(srv.ConfigFile); err != nil {
		return
	}
	if err = json.Unmarshal(contB, srv.Config); err != nil {
		return
	}
	return
}

func (srv *Server) PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (srv *Server) Run() (err error) {
	log.Println("Running REST Server")
	srv.apiEngine.Run(fmt.Sprintf("%s:%s", srv.Config.Server.Host, srv.Config.Server.Port))
	return
}
