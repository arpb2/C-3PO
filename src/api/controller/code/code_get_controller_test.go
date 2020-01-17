package code_test

import (
	"bytes"
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/arpb2/C-3PO/src/api/golden"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCodeGetControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", code.CreateGetController().Method)
}

func TestCodeGetControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id/codes/:code_id", code.CreateGetController().Path)
}

func TestCodeGetControllerBody_400OnEmptyUserId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
			Key:   "user_id",
			Value: "",
	})

	code.CreateGetController().Body(c)
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

	code.CreateGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.code_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodeGetControllerBody_500OnServiceReadError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	body := code.CreateGetBody(&SharedInMemoryCodeService{
		codeId: "1000",
		code:   nil,
		err:    errors.New("unexpected error"),
	})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1000",
	}, gin.Param{
		Key:   "code_id",
		Value: "1000",
	})

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_read.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodeGetControllerBody_400OnNoCodeStoredInService(t *testing.T) {
	gin.SetMode(gin.TestMode)

	body := code.CreateGetBody(&SharedInMemoryCodeService{
		codeId: "1000",
		code:   nil,
		err:    nil,
	})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1000",
	}, gin.Param{
		Key:   "code_id",
		Value: "1000",
	})

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.missing_code.read.service.golden.json")

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodeGetControllerBody_200OnCodeStoredOnService(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedCode := `
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello world!")
}
			`
	body := code.CreateGetBody(&SharedInMemoryCodeService{
		codeId: "1000",
		code:   &expectedCode,
		err:    nil,
	})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1000",
	}, gin.Param{
		Key:   "code_id",
		Value: "1000",
	})

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_code.golden.json")

	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodeGetControllerBody_200OnEmptyCodeStoredOnService(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedCode := ""
	body := code.CreateGetBody(&SharedInMemoryCodeService{
		codeId: "1000",
		code:   &expectedCode,
		err:    nil,
	})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1000",
	}, gin.Param{
		Key:   "code_id",
		Value: "1000",
	})

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_empty_code.golden.json")

	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.Equal(t, expected, actual)
}