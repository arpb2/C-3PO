package engine

import (
	"context"
	"errors"
	"fmt"
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ServerEngine interface {
	ControllerRegistrable
	http.Handler

	Run() error

	Shutdown() error
}

type ControllerRegistrable interface {

	Register(controller controller.Controller)

}

func CreateBasicServerEngine() ServerEngine {
	return defaultServerEngine{
		engine: gin.Default(),
		port:   GetPort(),
	}
}

var DefaultTokenHandler = jwt.TokenHandler{
	Secret: jwt.FetchJwtSecret(),
}

type defaultServerEngine struct {
	*http.Server
	engine *gin.Engine
	port   string
}

func (server defaultServerEngine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.engine.ServeHTTP(writer, request)
}

func (server defaultServerEngine) Run() error {
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

func (server defaultServerEngine) Shutdown() error {
	if server.Server == nil {
		return errors.New("no server running")
	}

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	log.Println("Server exiting")

	return nil
}

func (server defaultServerEngine) Register(controller controller.Controller) {
	var handlers []gin.HandlerFunc
	if controller.Middleware != nil {
		handlers = append(handlers, controller.Middleware...)
	}
	handlers = append(handlers, controller.Body)

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