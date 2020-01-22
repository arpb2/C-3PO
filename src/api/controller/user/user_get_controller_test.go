package user_test

import (
	"bytes"
	"errors"
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/executor/blocking"
	"github.com/arpb2/C-3PO/src/api/golden"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/http_wrapper/gin_wrapper"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/arpb2/C-3PO/src/api/service/user_service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func createGetController() controller.Controller {
	return user.CreateGetController(
		blocking.Executor{},
		single_auth.CreateMiddleware(
			jwt.CreateTokenHandler(),
		),
		user_service.CreateService(),
	)
}

func TestUserGetControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", createGetController().Method)
}

func TestUserGetControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id", createGetController().Path)
}

func TestUserGetControllerBody_400OnEmptyUserId(t *testing.T) {
	reader := new(http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("").Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	createGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserGetControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("not a number").Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	createGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserGetControllerBody_500OnServiceReadError(t *testing.T) {
	service := new(service.MockUserService)
	service.On("GetUser", uint(1000)).Return(nil, errors.New("whoops error")).Once()

	body := user.CreateGetBody(blocking.Executor{}, service)

	reader := new(http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_read.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}

func TestUserGetControllerBody_400OnNoUserStoredInService(t *testing.T) {
	service := new(service.MockUserService)
	service.On("GetUser", uint(1000)).Return(nil, nil).Once()

	body := user.CreateGetBody(blocking.Executor{}, service)

	reader := new(http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "not_found.missing_user.read.service.golden.json")

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}

func TestUserGetControllerBody_200OnUserStoredOnService(t *testing.T) {
	expectedUser := &model.User{
		Id:      1000,
		Email:   "test@email.com",
		Name:    "TestName",
		Surname: "TestSurname",
	}
	service := new(service.MockUserService)
	service.On("GetUser", uint(1000)).Return(expectedUser, nil).Once()

	body := user.CreateGetBody(blocking.Executor{}, service)

	reader := new(http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_user.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}