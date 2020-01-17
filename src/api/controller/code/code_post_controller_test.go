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

func TestCodePostControllerMethodIsPOST(t *testing.T) {
	assert.Equal(t, "POST", code.CreatePostController().Method)
}

func TestCodePostControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id/codes", code.CreatePostController().Path)
}

func TestCodePostControllerBody_400OnEmptyUserId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "",
	})

	code.CreatePostController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodePostControllerBody_400OnNoCode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1000",
	})
	c.Request, _ = http.NewRequest("GET", "/test", strings.NewReader(""))
	c.Request.PostForm = map[string][]string{}
	_ = c.Request.ParseForm()

	code.CreatePostController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.code.golden.json")

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodePostControllerBody_500OnServiceWriteError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	body := code.CreatePostBody(&SharedInMemoryCodeService{
		codeId: "1000",
		code:   nil,
		err:    errors.New("unexpected error"),
	})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
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

func TestCodePostControllerBody_200OnCodeStoredOnService(t *testing.T) {
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
	body := code.CreatePostBody(&SharedInMemoryCodeService{
		codeId: "1000",
		code:   nil,
		err:    nil,
	})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1000",
	})
	c.Request, _ = http.NewRequest("GET", "/test", strings.NewReader(""))
	c.Request.PostForm = map[string][]string{}
	c.Request.PostForm.Set("code", expectedCode)
	_ = c.Request.ParseForm()

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.write_code.golden.json")

	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.Equal(t, expected, actual)
}

func TestCodePostControllerBody_200OnEmptyCodeStoredOnService(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedCode := ""
	body := code.CreatePostBody(&SharedInMemoryCodeService{
		codeId: "1000",
		code:   nil,
		err:    nil,
	})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user_id",
		Value: "1000",
	})
	c.Request, _ = http.NewRequest("GET", "/test", strings.NewReader(""))
	c.Request.PostForm = map[string][]string{}
	c.Request.PostForm.Set("code", expectedCode)
	_ = c.Request.ParseForm()

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.write_empty_code.golden.json")

	assert.Equal(t, http.StatusOK, c.Writer.Status())
	assert.Equal(t, expected, actual)
}
