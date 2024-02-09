package gorpi

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

	"github.com/gauravsarma1992/go-rest-api/gorpi/middlewares"
	"github.com/gauravsarma1992/go-rest-api/gorpi/models"
	"github.com/gauravsarma1992/go-rest-api/gorpi/routing"
	"github.com/gin-gonic/gin"
)

var (
	ConfigFolderPath  = os.Getenv("CONFIG_FOLDER")
	DefaultConfigFile = ConfigFolderPath + "/config.json"
	DefaultHost       = "127.0.0.1"
	DefaulPort        = "9095"
)

type (
	Server struct {
		ctx    context.Context
		Cancel context.CancelFunc

		server    *http.Server
		apiEngine *gin.Engine

		RouteManager    *routing.RouteManager
		MiddlewareStack *middlewares.MiddlewareStack
		DB              *models.DB

		ConfigFile string
		Config     *Config

		closeCh chan bool
	}

	Config struct {
		Server struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"server"`
		Database struct {
			Username string `json:"username"`
			Password string `json:"password"`
			DbName   string `json:"db_name"`
			Host     string `json:"host"`
			Port     string `json:"port"`
		}
	}
)

func New(config *Config) (srv *Server, err error) {
	gin.SetMode(gin.ReleaseMode)
	srv = &Server{
		apiEngine: gin.Default(),
		Config:    config,
		closeCh:   make(chan bool),
	}
	if err = srv.Setup(); err != nil {
		return
	}
	return
}

func (srv *Server) Setup() (err error) {
	if srv.ctx, srv.Cancel = context.WithCancel(context.Background()); err != nil {
		return
	}
	if srv.Config == nil {
		if err = srv.setConfig(); err != nil {
			return
		}
	}
	if srv.DB, err = models.NewDB(
		srv.Config.Database.Username,
		srv.Config.Database.Password,
		srv.Config.Database.Host,
		srv.Config.Database.Port,
		srv.Config.Database.DbName,
	); err != nil {
		return
	}
	srv.MiddlewareStack = middlewares.NewMiddlewareStack(srv.DB)

	if err = srv.setRoutes(); err != nil {
		return
	}
	if err = srv.setHttpServer(); err != nil {
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

	// If the default config file doesn't exist, fallback to default constants
	if _, err = os.Stat(srv.ConfigFile); err != nil {
		srv.Config.Server.Host = DefaultHost
		srv.Config.Server.Port = DefaulPort
		err = nil
		return
	}

	if contB, err = ioutil.ReadFile(srv.ConfigFile); err != nil {
		return
	}
	if err = json.Unmarshal(contB, srv.Config); err != nil {
		return
	}
	return
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
