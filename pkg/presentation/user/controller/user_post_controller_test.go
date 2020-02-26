package controller_test

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
	debugpipeline "github.com/arpb2/C-3PO/test/mock/pipeline"

	usercontroller "github.com/arpb2/C-3PO/pkg/presentation/user/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/test/mock/golden"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createPostController() controller.Controller {
	return usercontroller.CreatePostController(debugpipeline.CreateDebugHttpPipeline(), nil, []validation.Validation{})
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

func TestUserPostControllerBody_500OnRepositoryCreateError(t *testing.T) {
	repository := new(repository.MockUserRepository)
	repository.On("CreateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(model2.User{}, errors.New("whoops error")).Once()

	body := usercontroller.CreatePostBody(debugpipeline.CreateDebugHttpPipeline(), repository, []validation.Validation{})

	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_write.repository.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestUserPostControllerBody_200OnUserStoredOnRepository(t *testing.T) {
	expectedUser := &model2.AuthenticatedUser{
		User: model2.User{
			Id:      1000,
			Email:   "test@email.com",
			Name:    "TestName",
			Surname: "TestSurname",
		},
		Password: "testpassword",
	}
	repository := new(repository.MockUserRepository)
	repository.On("CreateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(expectedUser.User, nil).Once()

	body := usercontroller.CreatePostBody(debugpipeline.CreateDebugHttpPipeline(), repository, []validation.Validation{})

	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model2.AuthenticatedUser) bool {
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
	repository.AssertExpectations(t)
}
