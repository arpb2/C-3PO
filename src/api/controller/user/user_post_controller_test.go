package user_test

import (
	"bytes"
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/golden"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/http_wrapper/gin_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service/user_service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func createPostController() controller.Controller {
	return user.CreatePostController(user_service.CreateService())
}

func TestUserPostControllerMethodIsPOST(t *testing.T) {
	assert.Equal(t, "POST", createPostController().Method)
}

func TestUserPostControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users", createPostController().Path)
}

func TestUserPostControllerBody_400OnEmptyOrMalformedUser(t *testing.T) {
	reader := new(http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(errors.New("malformed")).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	createPostController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserPostControllerBody_500OnServiceCreateError(t *testing.T) {
	service := new(MockUserService)
	service.On("CreateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(&model.AuthenticatedUser{}, errors.New("whoops error")).Once()

	body := user.CreatePostBody(service)

	reader := new(http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_write.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}

func TestUserPostControllerBody_500OnNoUserStoredInService(t *testing.T) {
	service := new(MockUserService)
	service.On("CreateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil, nil).Once()

	body := user.CreatePostBody(service)

	reader := new(http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_error.missing_user.write.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}

func TestUserPostControllerBody_200OnUserStoredOnService(t *testing.T) {
	expectedUser := &model.AuthenticatedUser{
		User: &model.User{
			Id:      1000,
			Email:   "test@email.com",
			Name:    "TestName",
			Surname: "TestSurname",
		},
		Password: "testpassword",
	}
	service := new(MockUserService)
	service.On("CreateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(expectedUser, nil).Once()

	body := user.CreatePostBody(service)

	reader := new(http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		return true
	})).Return(nil).Once()

	c, w := gin_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.create_user.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}
