package server_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/engine/c3po"
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
	err := server.StartApplication(&ServerEngineMock{})

	assert.Nil(t, err)
}

func TestStartApplicationFailureIsHandled(t *testing.T) {
	err := server.StartApplication(&FailingServerEngineMock{})

	assert.NotNil(t, err)
	assert.Equal(t, "woops this fails", err.Error())
}

func TestHealthGet(test *testing.T) {
	engine := c3po.New()
	server.RegisterRoutes(engine)

	w := performRequest(engine, "GET", "/ping", "")

	assert.Equal(test, 200, w.Code)
	assert.Equal(test, "pong", w.Body.String())
}

func TestCodeGetRegistered(test *testing.T) {
	engine := c3po.New()
	server.RegisterRoutes(engine)

	w := performRequest(engine, "GET", "/users/1/codes/1", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestCodePostRegistered(test *testing.T) {
	engine := c3po.New()
	server.RegisterRoutes(engine)

	w := performRequest(engine, "POST", "/users/1/codes", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestCodePutRegistered(test *testing.T) {
	engine := c3po.New()
	server.RegisterRoutes(engine)

	w := performRequest(engine, "PUT", "/users/1/codes/1", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestUserGetRegistered(test *testing.T) {
	engine := c3po.New()
	server.RegisterRoutes(engine)

	w := performRequest(engine, "GET", "/users/1", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestUserPostRegistered(test *testing.T) {
	engine := c3po.New()
	server.RegisterRoutes(engine)

	w := performRequest(engine, "POST", "/users", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestUserPutRegistered(test *testing.T) {
	engine := c3po.New()
	server.RegisterRoutes(engine)

	w := performRequest(engine, "PUT", "/users/1", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestUserDeleteRegistered(test *testing.T) {
	engine := c3po.New()
	server.RegisterRoutes(engine)

	w := performRequest(engine, "DELETE", "/users/1", "")

	assert.NotEqual(test, 404, w.Code)
}