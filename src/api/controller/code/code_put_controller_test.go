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
	"strings"
	"testing"
)

func TestCodePutControllerMethodIsPUT(t *testing.T) {
	assert.Equal(t, "PUT", code.PutController.Method)
}

func TestCodePutControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id/codes/:code_id", code.PutController.Path)
}

func TestCodePutControllerMiddleware_HasAuthenticationMiddleware(t *testing.T) {
	found := false

	for _, middleware := range code.PutController.Middleware {
		// Golang doesn't allow func comparisons, so we have to test identity through pointers using reflection.
		if reflect.ValueOf(auth.UserOrTeacherAuthenticationMiddleware).Pointer() == reflect.ValueOf(middleware).Pointer() {
			found = true
		}
	}

	assert.True(t, found)
}

func TestCodePutControllerBody_400OnEmptyUserId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "",
	})

	code.PutController.Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodePutControllerBody_400OnEmptyCodeId(t *testing.T) {
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

	code.PutController.Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.code_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodePutControllerBody_400OnNoCode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1000",
	}, gin.Param{
		Key:   "code_id",
		Value: "1000",
	})
	c.Request, _ = http.NewRequest("GET", "/test", strings.NewReader(""))
	c.Request.PostForm = map[string][]string{}
	_ = c.Request.ParseForm()

	code.PutController.Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.code.golden.json")

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodePutControllerBody_500OnServiceWriteError(t *testing.T) {
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
	c.Request, _ = http.NewRequest("GET", "/test", strings.NewReader(""))
	c.Request.PostForm = map[string][]string{}
	c.Request.PostForm.Set("code", "sending some code")
	_ = c.Request.ParseForm()

	code.PutController.Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_write.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodePutControllerBody_200OnCodeReplacedOnService(t *testing.T) {
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
	c.Request, _ = http.NewRequest("GET", "/test", strings.NewReader(""))
	c.Request.PostForm = map[string][]string{}
	c.Request.PostForm.Set("code", expectedCode)
	_ = c.Request.ParseForm()

	code.PutController.Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.replace_code.golden.json")

	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedCode, *code.Service.(*SharedInMemoryCodeService).code)
}

func TestCodePutControllerBody_200OnEmptyCodeStoredOnService(t *testing.T) {
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
	c.Request, _ = http.NewRequest("GET", "/test", strings.NewReader(""))
	c.Request.PostForm = map[string][]string{}
	c.Request.PostForm.Set("code", expectedCode)
	_ = c.Request.ParseForm()

	code.PutController.Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.replace_empty_code.golden.json")

	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedCode, *code.Service.(*SharedInMemoryCodeService).code)
}
