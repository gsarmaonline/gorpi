package gorestapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

var (
	ConfigFolderPath  = os.Getenv("CONFIG_FOLDER")
	DefaultConfigFile = ConfigFolderPath + "/config.json"
)

type (
	Server struct {
		ctx    context.Context
		Cancel context.CancelFunc

		server    *http.Server
		apiEngine *gin.Engine

		ConfigFile string
		Config     *Config

		closeCh chan bool
	}

	Config struct {
		Server struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"server"`
	}
)

func New(config *Config) (srv *Server, err error) {
	gin.SetMode(gin.ReleaseMode)
	srv = &Server{
		apiEngine: gin.Default(),
		Config:    config,
		closeCh:   make(chan bool),
	}
	if srv.ctx, srv.Cancel = context.WithCancel(context.Background()); err != nil {
		return
	}
	if srv.Config == nil {
		if err = srv.setConfig(); err != nil {
			return
		}
	}
	if err = srv.setHttpServer(); err != nil {
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

func (srv *Server) setHttpServer() (err error) {
	srv.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", srv.Config.Server.Host, srv.Config.Server.Port),
		Handler: srv.apiEngine,
	}
	return
}

func (srv *Server) handleShutdown() (err error) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-srv.closeCh:
			srv.Shutdown()
			return
		case <-quit:
			srv.Shutdown()
			return
		case <-srv.ctx.Done():
			srv.Shutdown()
			return
		}
	}
	return
}

func (srv *Server) Shutdown() (err error) {
	log.Println("Shutting down REST Server")
	if err = srv.server.Shutdown(srv.ctx); err != nil {
		log.Printf("Server Shutdown Failed:%+v", err)
	}
	return
}

func (srv *Server) Run() (err error) {
	log.Println("Running REST Server on", srv.Config.Server.Host, srv.Config.Server.Port)

	go func() {
		if err := srv.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	if err = srv.handleShutdown(); err != nil {
		return
	}

	return
}
