package controller_test

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	controller2 "github.com/arpb2/C-3PO/pkg/presentation/user/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/arpb2/C-3PO/pkg/infra/pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/arpb2/C-3PO/pkg/infra/executor"
	"github.com/arpb2/C-3PO/test/mock/golden"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createPostController() controller.Controller {
	return controller2.CreatePostController(pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()), nil, []validation.Validation{})
}

func TestUserPostControllerMethodIsPOST(t *testing.T) {
	assert.Equal(t, "POST", createPostController().Method)
}

func TestUserPostControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, "/users", createPostController().Path)
}

func TestUserPostControllerBody_400OnEmptyOrMalformedUser(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(errors.New("malformed")).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createPostController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserPostControllerBody_500OnServiceCreateError(t *testing.T) {
	service := new(service.MockUserService)
	service.On("CreateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(model.User{}, errors.New("whoops error")).Once()

	body := controller2.CreatePostBody(pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()), service, []validation.Validation{})

	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_write.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}

func TestUserPostControllerBody_200OnUserStoredOnService(t *testing.T) {
	expectedUser := &model.AuthenticatedUser{
		User: model.User{
			Id:      1000,
			Email:   "test@email.com",
			Name:    "TestName",
			Surname: "TestSurname",
		},
		Password: "testpassword",
	}
	service := new(service.MockUserService)
	service.On("CreateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(expectedUser.User, nil).Once()

	body := controller2.CreatePostBody(pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()), service, []validation.Validation{})

	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		return true
	})).Return(nil).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.create_user.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}
