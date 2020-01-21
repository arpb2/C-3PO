package code_test

import (
	"bytes"
	"errors"
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/arpb2/C-3PO/src/api/golden"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/http_wrapper/gin_wrapper"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/teacher_auth"
	"github.com/arpb2/C-3PO/src/api/service/code_service"
	"github.com/arpb2/C-3PO/src/api/service/teacher_service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func createPutController() controller.Controller {
	return code.CreatePutController(
		teacher_auth.CreateMiddleware(
			jwt.CreateTokenHandler(),
			teacher_service.GetService(),
		),
		code_service.GetService(),
	)
}

func TestCodePutControllerMethodIsPUT(t *testing.T) {
	assert.Equal(t, "PUT", createPutController().Method)
}

func TestCodePutControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id/codes/:code_id", createPutController().Path)
}

func TestCodePutControllerBody_400OnEmptyUserId(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("Param", "user_id").Return("").Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
}

func TestCodePutControllerBody_400OnEmptyCodeId(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("Param", "user_id").Return("1000").Once()
	reader.On("Param", "code_id").Return("").Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.code_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
}

func TestCodePutControllerBody_400OnNoCode(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("Param", "user_id").Return("1000").Once()
	reader.On("Param", "code_id").Return("1000").Once()
	reader.On("GetPostForm", "code").Return("", false).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.code.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
}

func TestCodePutControllerBody_500OnServiceWriteError(t *testing.T) {
	body := code.CreatePutBody(&SharedInMemoryCodeService{
		codeId: "1000",
		code:   nil,
		err:    errors.New("unexpected error"),
	})

	reader := new(http_wrapper.TestReader)
	reader.On("Param", "user_id").Return("1000").Once()
	reader.On("Param", "code_id").Return("1000").Once()
	reader.On("GetPostForm", "code").Return("sending some code", true).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_write.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
}

func TestCodePutControllerBody_200OnCodeReplacedOnService(t *testing.T) {
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

	reader := new(http_wrapper.TestReader)
	reader.On("Param", "user_id").Return("1000").Once()
	reader.On("Param", "code_id").Return("1000").Once()
	reader.On("GetPostForm", "code").Return(expectedCode, true).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.replace_code.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
}

func TestCodePutControllerBody_200OnEmptyCodeStoredOnService(t *testing.T) {
	expectedCode := ""
	body := code.CreatePutBody(&SharedInMemoryCodeService{
		codeId: "1000",
		code:   &expectedCode,
		err:    nil,
	})

	reader := new(http_wrapper.TestReader)
	reader.On("Param", "user_id").Return("1000").Once()
	reader.On("Param", "code_id").Return("1000").Once()
	reader.On("GetPostForm", "code").Return(expectedCode, true).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)


	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.replace_empty_code.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
}
