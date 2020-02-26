package server_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/arpb2/C-3PO/cmd/c3po/infrastructure/server"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/stretchr/testify/assert"
)

type MockServerEngine struct {
	mock.Mock
}

func (s *MockServerEngine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_ = s.Called(writer, request)
}
func (s *MockServerEngine) Run() error {
	args := s.Called()
	return args.Error(0)
}
func (s *MockServerEngine) Register(controller controller.Controller) {
	_ = s.Called(controller)
}

func TestStartApplicationSuccess(t *testing.T) {
	e := new(MockServerEngine)
	e.On("Run").Return(nil)
	err := server.StartApplication(e, []controller.Controller{})

	assert.Nil(t, err)
	e.AssertExpectations(t)
}

func TestStartApplicationFailureIsHandled(t *testing.T) {
	e := new(MockServerEngine)
	e.On("Run").Return(errors.New("woops this fails"))
	err := server.StartApplication(e, []controller.Controller{})

	assert.NotNil(t, err)
	assert.Equal(t, "woops this fails", err.Error())
	e.AssertExpectations(t)
}

func TestStartWithRegistration(t *testing.T) {
	e := new(MockServerEngine)
	e.On("Register", controller.Controller{}).Once()
	e.On("Run").Return(nil)
	err := server.StartApplication(e, []controller.Controller{
		{},
	})

	assert.Nil(t, err)
	e.AssertExpectations(t)
}
