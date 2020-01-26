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
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service/user_service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func createPutController() controller.Controller {
	return user.CreatePutController(
		single_auth.CreateMiddleware(
			jwt.CreateTokenHandler(),
		),
		user_service.CreateService(),
	)
}

func TestUserPutControllerMethodIsPUT(t *testing.T) {
	assert.Equal(t, "PUT", createPutController().Method)
}

func TestUserPutControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users/:user_id", createPutController().Path)
}

func TestUserPutControllerBody_400OnNoUserId(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("GetParameter", "user_id").Return("").Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserPutControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("GetParameter", "user_id").Return("not a number").Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserPutControllerBody_400OnEmptyOrMalformedUser(t *testing.T) {
	reader := new(http_wrapper.TestReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(errors.New("malformed")).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserPutControllerBody_500OnServiceCreateError(t *testing.T) {
	service := new(MockUserService)
	service.On("UpdateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(&model.User{
		Id:      1234,
		Email:   "test@email.com",
		Name:    "test name",
		Surname: "test surname",
	}, errors.New("whoops error")).Once()

	body := user.CreatePutBody(service)

	reader := new(http_wrapper.TestReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Run(func(args mock.Arguments) {
		args.Get(0).(*model.AuthenticatedUser).User = &model.User{
			Email:   "test@email.com",
			Name:    "test name",
			Surname: "test surname",
		}
		args.Get(0).(*model.AuthenticatedUser).Password = "test password"
	}).Return(nil).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_update.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}

func TestUserPutControllerBody_500OnNoUserStoredInService(t *testing.T) {
	service := new(MockUserService)
	service.On("UpdateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil, nil).Once()

	body := user.CreatePutBody(service)

	reader := new(http_wrapper.TestReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Run(func(args mock.Arguments) {
		args.Get(0).(*model.AuthenticatedUser).User = &model.User{
			Email:   "test@email.com",
			Name:    "test name",
			Surname: "test surname",
		}
		args.Get(0).(*model.AuthenticatedUser).Password = "test password"
	}).Return(nil).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_error.missing_user.update.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}

func TestUserPutControllerBody_400OnIdSpecified(t *testing.T) {
	service := new(MockUserService)
	body := user.CreatePutBody(service)

	reader := new(http_wrapper.TestReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		return true
	})).Run(func(args mock.Arguments) {
		args.Get(0).(*model.AuthenticatedUser).User = &model.User{
			Id: 1000,
		}
	}).Return(nil).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.update_user.id_specified.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}

func TestUserPutControllerBody_200OnUserStoredOnService(t *testing.T) {
	expectedUser := &model.User{
		Id:      1000,
		Email:   "test@email.com",
		Name:    "TestName",
		Surname: "TestSurname",
	}
	service := new(MockUserService)
	service.On("UpdateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(expectedUser, nil).Once()

	body := user.CreatePutBody(service)

	reader := new(http_wrapper.TestReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		return true
	})).Run(func(args mock.Arguments) {
		args.Get(0).(*model.AuthenticatedUser).User = &model.User{
			Email:   "test@email.com",
			Name:    "test name",
			Surname: "test surname",
		}
		args.Get(0).(*model.AuthenticatedUser).Password = "test password"
	}).Return(nil).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.update_user.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}