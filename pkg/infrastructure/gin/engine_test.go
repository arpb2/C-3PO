package gin_test

import (
	"net/http/httptest"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/infrastructure/gin"
	"github.com/stretchr/testify/mock"
)

type mockHandler struct {
	mock.Mock
}

func (h *mockHandler) Handle(ctx *http.Context) {
	_ = h.Called(ctx)
}

func TestEngine_GivenAGet_WhenServing_ThenItsAvailable(t *testing.T) {
	handler := new(mockHandler)
	handler.On("Handle", mock.Anything).Once()
	e := gin.CreateEngine()
	e.GET("/test", handler.Handle)
	r := httptest.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, r)

	handler.AssertExpectations(t)
}

func TestEngine_GivenAPost_WhenServing_ThenItsAvailable(t *testing.T) {
	handler := new(mockHandler)
	handler.On("Handle", mock.Anything).Once()
	e := gin.CreateEngine()
	e.POST("/test", handler.Handle)
	r := httptest.NewRequest("POST", "/test", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, r)

	handler.AssertExpectations(t)
}

func TestEngine_GivenAPut_WhenServing_ThenItsAvailable(t *testing.T) {
	handler := new(mockHandler)
	handler.On("Handle", mock.Anything).Once()
	e := gin.CreateEngine()
	e.PUT("/test", handler.Handle)
	r := httptest.NewRequest("PUT", "/test", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, r)

	handler.AssertExpectations(t)
}

func TestEngine_GivenADelete_WhenServing_ThenItsAvailable(t *testing.T) {
	handler := new(mockHandler)
	handler.On("Handle", mock.Anything).Once()
	e := gin.CreateEngine()
	e.DELETE("/test", handler.Handle)
	r := httptest.NewRequest("DELETE", "/test", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, r)

	handler.AssertExpectations(t)
}

func TestEngine_GivenAPatch_WhenServing_ThenItsAvailable(t *testing.T) {
	handler := new(mockHandler)
	handler.On("Handle", mock.Anything).Once()
	e := gin.CreateEngine()
	e.PATCH("/test", handler.Handle)
	r := httptest.NewRequest("PATCH", "/test", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, r)

	handler.AssertExpectations(t)
}

func TestEngine_GivenAGlobalMiddleware_WhenServing_ThenItsUsed(t *testing.T) {
	handler := new(mockHandler)
	handler.On("Handle", mock.Anything).Twice()
	e := gin.CreateEngine()
	e.Use(handler.Handle)
	e.GET("/test", handler.Handle)
	r := httptest.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, r)

	handler.AssertExpectations(t)
}
