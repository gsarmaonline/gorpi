package gorestapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	ConfigFolderPath  = os.Getenv("CONFIG_FOLDER")
	DefaultConfigFile = ConfigFolderPath + "/config.json"
)

type (
	Server struct {
		apiEngine  *gin.Engine
		ConfigFile string
		Config     *Config
		Generator  *Generator
	}

	Config struct {
		Server struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"server"`
	}
)

func New(config *Config) (srv *Server, err error) {
	srv = &Server{
		apiEngine: gin.Default(),
		Config:    config,
	}
	if srv.Config == nil {
		if err = srv.setConfig(); err != nil {
			return
		}
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
	srv.Config = &Config{}
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

func (srv *Server) generateRestApis(object interface{}) (err error) {
	return
}
