package code_test

import (
	"bytes"
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/arpb2/C-3PO/src/api/golden"
	"github.com/arpb2/C-3PO/src/api/middleware/auth"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCodeGetControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", code.GetController.Method)
}

func TestCodeGetControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id/codes/:code_id", code.GetController.Path)
}

func TestCodeGetControllerMiddleware_HasAuthenticationMiddleware(t *testing.T) {
	found := false

	for _, middleware := range code.GetController.Middleware {
		// Golang doesn't allow func comparisons, so we have to test identity through pointers using reflection.
		if reflect.ValueOf(auth.UserOrTeacherAuthenticationMiddleware).Pointer() == reflect.ValueOf(middleware).Pointer() {
			found = true
		}
	}

	assert.True(t, found)
}

func TestCodeGetControllerBody_400OnEmptyUserId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
			Key:   "user_id",
			Value: "",
	})

	code.GetController.Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodeGetControllerBody_400OnEmptyCodeId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
			Key:   "user_id",
			Value: "1000",
	}, gin.Param{
			Key:   "code_id",
			Value: "",
	})

	code.GetController.Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.code_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodeGetControllerBody_500OnServiceReadError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	defer func() {
		code.Service = nil
	}()
	code.Service = &SharedInMemoryCodeService{
		codeId: "1000",
		code:   nil,
		err:    errors.New("Unexpected error"),
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1000",
	}, gin.Param{
		Key:   "code_id",
		Value: "1000",
	})

	code.GetController.Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_read.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodeGetControllerBody_400OnNoCodeStoredInService(t *testing.T) {
	gin.SetMode(gin.TestMode)

	defer func() {
		code.Service = nil
	}()
	code.Service = &SharedInMemoryCodeService{
		codeId: "1000",
		code:   nil,
		err:    nil,
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1000",
	}, gin.Param{
		Key:   "code_id",
		Value: "1000",
	})

	code.GetController.Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.missing_code.read.service.golden.json")

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodeGetControllerBody_200OnCodeStoredOnService(t *testing.T) {
	gin.SetMode(gin.TestMode)

	defer func() {
		code.Service = nil
	}()
	expectedCode := `
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello world!")
}
			`
	code.Service = &SharedInMemoryCodeService{
		codeId: "1000",
		code:   &expectedCode,
		err:    nil,
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1000",
	}, gin.Param{
		Key:   "code_id",
		Value: "1000",
	})

	code.GetController.Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_code.golden.json")

	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodeGetControllerBody_200OnEmptyCodeStoredOnService(t *testing.T) {
	gin.SetMode(gin.TestMode)

	defer func() {
		code.Service = nil
	}()
	expectedCode := ""
	code.Service = &SharedInMemoryCodeService{
		codeId: "1000",
		code:   &expectedCode,
		err:    nil,
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1000",
	}, gin.Param{
		Key:   "code_id",
		Value: "1000",
	})

	code.GetController.Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_empty_code.golden.json")

	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.Equal(t, expected, actual)
}