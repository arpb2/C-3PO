package gin

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/engine"
	ginwrapper "github.com/arpb2/C-3PO/pkg/infra/http/gin"
	"github.com/gin-gonic/gin"
)

func CreateEngine() engine.ServerEngine {
	return &serverEngine{
		engine: gin.Default(),
		port:   GetPort(),
	}
}

type serverEngine struct {
	*http.Server
	engine *gin.Engine
	port   string
}

func (server serverEngine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.engine.ServeHTTP(writer, request)
}

func (server serverEngine) Run() error {
	if server.Server != nil {
		return errors.New("can't ignite, server already running")
	}

	addr := ":" + server.port
	fmt.Printf("Listening and serving HTTP on %s\n", addr)

	server.Server = &http.Server{
		Addr:    addr,
		Handler: server.engine,
	}

	return server.ListenAndServe()
}

func (server serverEngine) Register(controller controller.Controller) {
	var handlers []gin.HandlerFunc
	if controller.Middleware != nil {
		handlers = append(handlers, ginwrapper.CreateHandlers(controller.Middleware...)...)
	}
	handlers = append(handlers, ginwrapper.CreateHandler(controller.Body))

	server.engine.Handle(
		controller.Method,
		controller.Path,
		handlers...,
	)
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
