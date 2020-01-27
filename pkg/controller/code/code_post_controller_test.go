package code_controller_test

import (
	"bytes"
	"errors"
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/auth/jwt"
	"github.com/arpb2/C-3PO/pkg/controller/code"
	"github.com/arpb2/C-3PO/pkg/executor"
	"github.com/arpb2/C-3PO/pkg/middleware/auth/teacher_auth"
	code_service "github.com/arpb2/C-3PO/pkg/service/code"
	"github.com/arpb2/C-3PO/pkg/service/teacher"
	"github.com/arpb2/C-3PO/pkg/service/user"
	"github.com/arpb2/C-3PO/hack/golden"
	test_http_wrapper "github.com/arpb2/C-3PO/hack/http_wrapper"
	"github.com/arpb2/C-3PO/hack/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func createPostController() controller.Controller {
	return code_controller.CreatePostController(
		executor.CreatePipeline(executor.CreateDebugHttpExecutor()),
		teacher_auth.CreateMiddleware(
			jwt.CreateTokenHandler(),
			teacher_service.CreateService(user_service.CreateService()),
		),
		code_service.CreateService(),
	)
}

func TestCodePostControllerMethodIsPOST(t *testing.T) {
	assert.Equal(t, "POST", createPostController().Method)
}

func TestCodePostControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id/codes", createPostController().Path)
}

func TestCodePostControllerBody_400OnEmptyUserId(t *testing.T) {
	reader := new(test_http_wrapper.MockReader)
	reader.On("GetFormData", "code").Return("", true).Maybe()
	reader.On("GetParameter", "user_id").Return("").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	createPostController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestCodePostControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(test_http_wrapper.MockReader)
	reader.On("GetFormData", "code").Return("", true).Maybe()
	reader.On("GetParameter", "user_id").Return("not a number").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	createPostController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestCodePostControllerBody_400OnNoCode(t *testing.T) {
	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetFormData", "code").Return("", false).Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	createPostController().Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.code.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestCodePostControllerBody_500OnServiceWriteError(t *testing.T) {
	codeService := new(service.MockCodeService)
	codeService.On("GetCode", uint(1000), uint(1000)).Return(nil, errors.New("whoops error"))

	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "code_id").Return("1000").Maybe()
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetFormData", "code").Return("sending some code", true).Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	code_controller.CreateGetBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), codeService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_write.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	codeService.AssertExpectations(t)
}

func TestCodePostControllerBody_200OnCodeStoredOnService(t *testing.T) {
	expectedCode := `
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello world!")
}
			`
	codeService := new(service.MockCodeService)
	codeService.On("GetCode", uint(1000), uint(1000)).Return(&model.Code{
		UserId: 1000,
		Id:     1000,
		Code:   expectedCode,
	}, nil)

	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "code_id").Return("1000").Maybe()
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetFormData", "code").Return(expectedCode, true).Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	code_controller.CreateGetBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), codeService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.write_code.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	codeService.AssertExpectations(t)
}

func TestCodePostControllerBody_200OnEmptyCodeStoredOnService(t *testing.T) {
	expectedCode := ""
	codeService := new(service.MockCodeService)
	codeService.On("GetCode", uint(1000), uint(1000)).Return(&model.Code{
		UserId: 1000,
		Id:     1000,
		Code:   expectedCode,
	}, nil)

	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "code_id").Return("1000").Maybe()
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetFormData", "code").Return(expectedCode, true).Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	code_controller.CreateGetBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), codeService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.write_empty_code.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	codeService.AssertExpectations(t)
}
