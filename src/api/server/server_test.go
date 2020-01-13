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
func (server ServerEngineMock) Shutdown() error {
	return nil
}

type FailingServerEngineMock struct{}

func (server FailingServerEngineMock) ServeHTTP(writer http.ResponseWriter, request *http.Request) {}
func (server FailingServerEngineMock) Run() error {
	return errors.New("woops this fails")
}
func (server FailingServerEngineMock) Register(controller controller.Controller) {}
func (server FailingServerEngineMock) Shutdown() error {
	return nil
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

func TestHealthGet(test *testing.T) {
	w := performRequest(server.CreateEngine(), "GET", "/ping", "")

	assert.Equal(test, 200, w.Code)
	assert.Equal(test, "pong", w.Body.String())
}

func TestCodeGetRegistered(test *testing.T) {
	w := performRequest(server.CreateEngine(), "GET", "/users/1/codes/1", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestCodePostRegistered(test *testing.T) {
	w := performRequest(server.CreateEngine(), "POST", "/users/1/codes", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestCodePutRegistered(test *testing.T) {
	w := performRequest(server.CreateEngine(), "PUT", "/users/1/codes/1", "")

	assert.NotEqual(test, 404, w.Code)
}