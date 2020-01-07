package server_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/server"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func performRequest(r http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

type ServerEngineMock struct{}

func (server ServerEngineMock) ServeHTTP(writer http.ResponseWriter, request *http.Request) {}
func (server ServerEngineMock) Run() error {
	return nil
}
func (server ServerEngineMock) Register(controller controller.Controller) {}

type FailingServerEngineMock struct{}

func (server FailingServerEngineMock) ServeHTTP(writer http.ResponseWriter, request *http.Request) {}
func (server FailingServerEngineMock) Run() error {
	return errors.New("woops this fails")
}
func (server FailingServerEngineMock) Register(controller controller.Controller) {}

func TestHealthGet(test *testing.T) {
	w := performRequest(server.CreateEngine(), "GET", "/ping", "")

	assert.Equal(test, 200, w.Code)
	assert.Equal(test, "pong", w.Body.String())
	server.Engine = server.CreateEngine()
}

func TestStartApplicationSuccess(t *testing.T) {
	server.Engine = ServerEngineMock{}
	err := server.StartApplication()

	assert.Nil(t, err)
	server.Engine = server.CreateEngine()
}

func TestStartApplicationFailureIsHandled(t *testing.T) {
	server.Engine = FailingServerEngineMock{}
	err := server.StartApplication()

	assert.NotNil(t, err)
	assert.Equal(t, "woops this fails", err.Error())
	server.Engine = server.CreateEngine()
}