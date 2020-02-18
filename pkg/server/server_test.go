package server_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/arpb2/C-3PO/api/controller"
	ginengine "github.com/arpb2/C-3PO/pkg/engine/gin"
	"github.com/arpb2/C-3PO/pkg/server"
	"github.com/stretchr/testify/assert"
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
	os.Setenv("MYSQL_DSN", "root:password@tcp(localhost:3306)/db")
	defer os.Unsetenv("MYSQL_DSN")
	err := server.StartApplication(&ServerEngineMock{})

	assert.Nil(t, err)
}

func TestStartApplicationFailureIsHandled(t *testing.T) {
	os.Setenv("MYSQL_DSN", "root:password@tcp(localhost:3306)/db")
	defer os.Unsetenv("MYSQL_DSN")
	err := server.StartApplication(&FailingServerEngineMock{})

	assert.NotNil(t, err)
	assert.Equal(t, "woops this fails", err.Error())
}

func TestHealthGet(test *testing.T) {
	os.Setenv("MYSQL_DSN", "root:password@tcp(localhost:3306)/db")
	defer os.Unsetenv("MYSQL_DSN")

	engine := ginengine.New()
	binders, deferClose := server.CreateBinders()
	defer deferClose()
	server.RegisterRoutes(engine, binders)

	w := performRequest(engine, "GET", "/ping", "")

	assert.Equal(test, 200, w.Code)
	assert.Equal(test, "pong", w.Body.String())
}

func TestUserLevelGetRegistered(test *testing.T) {
	os.Setenv("MYSQL_DSN", "root:password@tcp(localhost:3306)/db")
	defer os.Unsetenv("MYSQL_DSN")

	engine := ginengine.New()
	binders, deferClose := server.CreateBinders()
	defer deferClose()
	server.RegisterRoutes(engine, binders)

	w := performRequest(engine, "GET", "/users/1/levels/1", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestUserLevelPutRegistered(test *testing.T) {
	os.Setenv("MYSQL_DSN", "root:password@tcp(localhost:3306)/db")
	defer os.Unsetenv("MYSQL_DSN")

	engine := ginengine.New()
	binders, deferClose := server.CreateBinders()
	defer deferClose()
	server.RegisterRoutes(engine, binders)

	w := performRequest(engine, "PUT", "/users/1/levels/1", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestUserGetRegistered(test *testing.T) {
	os.Setenv("MYSQL_DSN", "root:password@tcp(localhost:3306)/db")
	defer os.Unsetenv("MYSQL_DSN")

	engine := ginengine.New()
	binders, deferClose := server.CreateBinders()
	defer deferClose()
	server.RegisterRoutes(engine, binders)

	w := performRequest(engine, "GET", "/users/1", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestUserPostRegistered(test *testing.T) {
	os.Setenv("MYSQL_DSN", "root:password@tcp(localhost:3306)/db")
	defer os.Unsetenv("MYSQL_DSN")

	engine := ginengine.New()
	binders, deferClose := server.CreateBinders()
	defer deferClose()
	server.RegisterRoutes(engine, binders)

	w := performRequest(engine, "POST", "/users", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestUserPutRegistered(test *testing.T) {
	os.Setenv("MYSQL_DSN", "root:password@tcp(localhost:3306)/db")
	defer os.Unsetenv("MYSQL_DSN")

	engine := ginengine.New()
	binders, deferClose := server.CreateBinders()
	defer deferClose()
	server.RegisterRoutes(engine, binders)

	w := performRequest(engine, "PUT", "/users/1", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestUserDeleteRegistered(test *testing.T) {
	os.Setenv("MYSQL_DSN", "root:password@tcp(localhost:3306)/db")
	defer os.Unsetenv("MYSQL_DSN")

	engine := ginengine.New()
	binders, deferClose := server.CreateBinders()
	defer deferClose()
	server.RegisterRoutes(engine, binders)

	w := performRequest(engine, "DELETE", "/users/1", "")

	assert.NotEqual(test, 404, w.Code)
}

func TestSessionPostRegistered(test *testing.T) {
	os.Setenv("MYSQL_DSN", "root:password@tcp(localhost:3306)/db")
	defer os.Unsetenv("MYSQL_DSN")

	engine := ginengine.New()
	binders, deferClose := server.CreateBinders()
	defer deferClose()
	server.RegisterRoutes(engine, binders)

	w := performRequest(engine, "POST", "/session", "")

	assert.NotEqual(test, 404, w.Code)
}
