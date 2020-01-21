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
	"strings"
	"testing"
)

func TestCodePutControllerMethodIsPUT(t *testing.T) {
	assert.Equal(t, "PUT", code.CreatePutController().Method)
}

func TestCodePutControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id/codes/:code_id", code.CreatePutController().Path)
}

func TestCodePutControllerBody_400OnEmptyUserId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "",
	})

	code.CreatePutController().Body(c)
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

	code.CreatePutController().Body(c)
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

	code.CreatePutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.code.golden.json")

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodePutControllerBody_500OnServiceWriteError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	body := code.CreatePutBody(&SharedInMemoryCodeService{
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
	c.Request, _ = http.NewRequest("GET", "/test", strings.NewReader(""))
	c.Request.PostForm = map[string][]string{}
	c.Request.PostForm.Set("code", "sending some code")
	_ = c.Request.ParseForm()

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_write.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodePutControllerBody_200OnCodeReplacedOnService(t *testing.T) {
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
	body := code.CreatePutBody(&SharedInMemoryCodeService{
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
	c.Request, _ = http.NewRequest("GET", "/test", strings.NewReader(""))
	c.Request.PostForm = map[string][]string{}
	c.Request.PostForm.Set("code", expectedCode)
	_ = c.Request.ParseForm()

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.replace_code.golden.json")

	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodePutControllerBody_200OnEmptyCodeStoredOnService(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedCode := ""
	body := code.CreatePutBody(&SharedInMemoryCodeService{
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
	c.Request, _ = http.NewRequest("GET", "/test", strings.NewReader(""))
	c.Request.PostForm = map[string][]string{}
	c.Request.PostForm.Set("code", expectedCode)
	_ = c.Request.ParseForm()

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.replace_empty_code.golden.json")

	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.Equal(t, expected, actual)
}
