package controller_test

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHalt_OnResponseSuccess_DoNothing(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = &http.Request{
		URL: &url.URL{
			Scheme: "https",
			Host:   "localhost",
			Path:   "/test",
		},
	}

	controller.Halt(c, http.StatusCreated, "no error, success")

	assert.NotEqual(t, http.StatusCreated, recorder.Code)
	assert.Zero(t, recorder.Body.Len())
}

func TestHalt_OnResponseFailure_SetsErrorJson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = &http.Request{
		URL: &url.URL{
			Scheme: "https",
			Host:   "localhost",
			Path:   "/test",
		},
	}

	errorMessage := "some error"

	controller.Halt(c, http.StatusUnauthorized, errorMessage)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Contains(t, recorder.Body.String(), errorMessage)
}
