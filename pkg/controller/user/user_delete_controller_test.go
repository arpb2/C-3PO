package user_controller_test

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/hack/golden"
	test_http_wrapper "github.com/arpb2/C-3PO/hack/http_wrapper"
	"github.com/arpb2/C-3PO/hack/service"
	"github.com/arpb2/C-3PO/pkg/auth/jwt"
	user_controller "github.com/arpb2/C-3PO/pkg/controller/user"
	"github.com/arpb2/C-3PO/pkg/executor"
	"github.com/arpb2/C-3PO/pkg/middleware/auth/single_auth"
	user_service "github.com/arpb2/C-3PO/pkg/service/user"
	"github.com/stretchr/testify/assert"
)

func createDeleteController() controller.Controller {
	return user_controller.CreateDeleteController(
		executor.CreatePipeline(executor.CreateDebugHttpExecutor()),
		single_auth.CreateMiddleware(
			jwt.CreateTokenHandler(),
		),
		user_service.CreateService(),
	)
}

func TestUserDeleteControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "DELETE", createDeleteController().Method)
}

func TestUserDeleteControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id", createDeleteController().Path)
}

func TestUserDeleteControllerBody_400OnNoUserId(t *testing.T) {
	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	createDeleteController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserDeleteControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("not a number").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	createDeleteController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserDeleteControllerBody_500OnServiceDeleteError(t *testing.T) {
	service := new(service.MockUserService)
	service.On("DeleteUser", uint(1000)).Return(errors.New("whoops error")).Once()

	body := user_controller.CreateDeleteBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), service)

	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_delete.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}

func TestUserDeleteControllerBody_200OnUserDeletedOnService(t *testing.T) {
	service := new(service.MockUserService)
	service.On("DeleteUser", uint(1000)).Return(nil).Once()

	body := user_controller.CreateDeleteBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), service)

	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Zero(t, w.Body.Len())
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}
