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

func createGetController() controller.Controller {
	return code_controller.CreateGetController(
		executor.CreatePipeline(executor.CreateDebugHttpExecutor()),
		teacher_auth.CreateMiddleware(
			jwt.CreateTokenHandler(),
			teacher_service.CreateService(user_service.CreateService()),
		),
		code_service.CreateService(),
	)
}

func TestCodeGetControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", createGetController().Method)
}

func TestCodeGetControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id/codes/:code_id", createGetController().Path)
}

func TestCodeGetControllerBody_400OnEmptyUserId(t *testing.T) {
	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "code_id").Return("1").Maybe()
	reader.On("GetParameter", "user_id").Return("").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	createGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestCodeGetControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "code_id").Return("1000").Maybe()
	reader.On("GetParameter", "user_id").Return("not a number").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	createGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestCodeGetControllerBody_400OnMalformedCodeId(t *testing.T) {
	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "code_id").Return("not a number").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	createGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.code_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestCodeGetControllerBody_400OnEmptyCodeId(t *testing.T) {
	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "code_id").Return("").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	createGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.code_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestCodeGetControllerBody_500OnServiceReadError(t *testing.T) {
	codeService := new(service.MockCodeService)
	codeService.On("GetCode", uint(1000), uint(1000)).Return(nil, errors.New("whoops error"))

	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "code_id").Return("1000").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	code_controller.CreateGetBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), codeService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_read.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	codeService.AssertExpectations(t)
}

func TestCodeGetControllerBody_400OnNoCodeStoredInService(t *testing.T) {
	codeService := new(service.MockCodeService)
	codeService.On("GetCode", uint(1000), uint(1000)).Return(nil, nil)

	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "code_id").Return("1000").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	code_controller.CreateGetBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), codeService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "not_found.missing_code.read.service.golden.json")

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	codeService.AssertExpectations(t)
}

func TestCodeGetControllerBody_200OnCodeStoredOnService(t *testing.T) {
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
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "code_id").Return("1000").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	code_controller.CreateGetBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), codeService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_code.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	codeService.AssertExpectations(t)
}

func TestCodeGetControllerBody_200OnEmptyCodeStoredOnService(t *testing.T) {
	expectedCode := ""
	codeService := new(service.MockCodeService)
	codeService.On("GetCode", uint(1000), uint(1000)).Return(&model.Code{
		UserId: 1000,
		Id:     1000,
		Code:   expectedCode,
	}, nil)

	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "code_id").Return("1000").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	code_controller.CreateGetBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), codeService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_empty_code.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	codeService.AssertExpectations(t)
}
