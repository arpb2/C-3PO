package gin

import (
	"errors"
	"fmt"
	"net/http"

	http2 "github.com/arpb2/C-3PO/pkg/domain/http"

	"github.com/arpb2/C-3PO/pkg/domain/engine"
	"github.com/gin-gonic/gin"
)

func CreateEngine() engine.ServerEngine {
	return &serverEngine{
		engine: gin.Default(),
	}
}

type serverEngine struct {
	*http.Server
	engine *gin.Engine
}

func (s *serverEngine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.engine.ServeHTTP(writer, request)
}

func (s *serverEngine) Run(port string) error {
	if s.Server != nil {
		return errors.New("can't ignite, server already running")
	}

	addr := ":" + port
	fmt.Printf("Listening and serving HTTP on %s\n", addr)

	s.Server = &http.Server{
		Addr:    addr,
		Handler: s.engine,
	}

	return s.ListenAndServe()
}

func (s *serverEngine) GET(url string, handlers ...http2.Handler) {
	s.engine.GET(url, CreateHandlers(handlers...)...)
}

func (s *serverEngine) POST(url string, handlers ...http2.Handler) {
	s.engine.POST(url, CreateHandlers(handlers...)...)
}

func (s *serverEngine) PUT(url string, handlers ...http2.Handler) {
	s.engine.PUT(url, CreateHandlers(handlers...)...)
}

func (s *serverEngine) PATCH(url string, handlers ...http2.Handler) {
	s.engine.PATCH(url, CreateHandlers(handlers...)...)
}

func (s *serverEngine) DELETE(url string, handlers ...http2.Handler) {
	s.engine.DELETE(url, CreateHandlers(handlers...)...)
}

func (s *serverEngine) Use(handlers ...http2.Handler) {
	s.engine.Use(CreateHandlers(handlers...)...)
}
