package engine

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type ServerEngine interface {
	ServeHTTP(writer http.ResponseWriter, request *http.Request)

	Register(controller controller.Controller)

	Run() error
}

func CreateBasicServerEngine() ServerEngine {
	return defaultServerEngine{
		engine: gin.Default(),
		port:   GetPort(),
	}
}

type defaultServerEngine struct {
	engine *gin.Engine
	port   string
}

func (server defaultServerEngine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.engine.ServeHTTP(writer, request)
}

func (server defaultServerEngine) Run() error {
	return server.engine.Run(":" + server.port)
}

func (server defaultServerEngine) Register(controller controller.Controller) {
	server.engine.Handle(
		controller.Method,
		controller.Path,
		controller.Body,
	)
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}