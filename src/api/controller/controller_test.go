package controller_test

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/http_wrapper/gin_wrapper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHalt_OnResponseSuccess_DoNothing(t *testing.T) {
	reader := new(http_wrapper.MockReader)
	reader.On("GetUrl").Return("https://localhost/test").Once()

	c, recorder := gin_wrapper.CreateTestContext()
	c.Reader = reader

	controller.Halt(c, http.StatusCreated, "no error, success")

	assert.NotEqual(t, http.StatusCreated, recorder.Code)
	assert.Zero(t, recorder.Body.Len())
	reader.AssertExpectations(t)
}

func TestHalt_OnResponseFailure_SetsErrorJson(t *testing.T) {
	c, recorder := gin_wrapper.CreateTestContext()

	errorMessage := "some error"

	controller.Halt(c, http.StatusUnauthorized, errorMessage)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Contains(t, recorder.Body.String(), errorMessage)
}
