package user_test

import (
	"bytes"
	"errors"
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/golden"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/http_wrapper/gin_wrapper"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/arpb2/C-3PO/src/api/service/user_service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func createDeleteController() controller.Controller {
	return user.CreateDeleteController(
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
	reader := new(http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("").Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	createDeleteController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserDeleteControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("not a number").Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	createDeleteController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserDeleteControllerBody_500OnServiceDeleteError(t *testing.T) {
	service := new(MockUserService)
	service.On("DeleteUser", uint(1000)).Return(errors.New("whoops error")).Once()

	body := user.CreateDeleteBody(service)

	reader := new(http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()

	c, w := gin_wrapper.CreateTestContext()
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
	service := new(MockUserService)
	service.On("DeleteUser", uint(1000)).Return(nil).Once()

	body := user.CreateDeleteBody(service)

	reader := new(http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Zero(t, w.Body.Len())
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}