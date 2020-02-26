package gin

import (
	"errors"
	"fmt"
	"net/http"

	ginwrapper "github.com/arpb2/C-3PO/cmd/c3po/infrastructure/http/gin"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/engine"
	"github.com/gin-gonic/gin"
)

func CreateEngine(port string) engine.ServerEngine {
	return &serverEngine{
		engine: gin.Default(),
		port:   port,
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
