package controller_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/user"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"
	"github.com/arpb2/C-3PO/test/mock/token"

	usercontroller "github.com/arpb2/C-3PO/pkg/presentation/user/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/auth/user/single"
	"github.com/arpb2/C-3PO/test/mock/golden"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createPutController() controller.Controller {
	return usercontroller.CreatePutController(
		pipeline2.CreateDebugHttpPipeline(),
		single.CreateMiddleware(
			&token.MockTokenHandler{},
		),
		nil,
		[]validation.Validation{},
	)
}

func TestUserPutControllerMethodIsPUT(t *testing.T) {
	assert.Equal(t, "PUT", createPutController().Method)
}

func TestUserPutControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("/users/:%s", user.ParamUserId), createPutController().Path)
}

func TestUserPutControllerBody_400OnNoUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil).Maybe()
	reader.On("GetParameter", user.ParamUserId).Return("").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserPutControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil).Maybe()
	reader.On("GetParameter", user.ParamUserId).Return("not a number").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserPutControllerBody_400OnEmptyOrMalformedUser(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(errors.New("malformed")).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserPutControllerBody_500OnRepositoryCreateError(t *testing.T) {
	repository := new(repository.MockUserRepository)
	repository.On("UpdateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(model2.User{
		Id:      1234,
		Email:   "test@email.com",
		Name:    "test name",
		Surname: "test surname",
	}, errors.New("whoops error")).Once()

	body := usercontroller.CreatePutBody(pipeline2.CreateDebugHttpPipeline(), repository, []validation.Validation{})

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Run(func(args mock.Arguments) {
		args.Get(0).(*model2.AuthenticatedUser).User = model2.User{
			Email:   "test@email.com",
			Name:    "test name",
			Surname: "test surname",
		}
		args.Get(0).(*model2.AuthenticatedUser).Password = "test password"
	}).Return(nil).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_update.repository.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestUserPutControllerBody_400OnIdSpecified(t *testing.T) {
	repository := new(repository.MockUserRepository)

	body := usercontroller.CreatePutBody(pipeline2.CreateDebugHttpPipeline(), repository, []validation.Validation{
		validation.IdProvided,
	})

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()
	reader.On("ReadBody", mock.MatchedBy(func(obj *model2.AuthenticatedUser) bool {
		obj.User = model2.User{
			Id: 1000,
		}
		return true
	})).Return(nil).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.update_user.id_specified.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestUserPutControllerBody_200OnUserStoredOnRepository(t *testing.T) {
	expectedUser := model2.User{
		Id:      1000,
		Email:   "test@email.com",
		Name:    "TestName",
		Surname: "TestSurname",
	}
	repository := new(repository.MockUserRepository)
	repository.On("UpdateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(expectedUser, nil).Once()

	body := usercontroller.CreatePutBody(pipeline2.CreateDebugHttpPipeline(), repository, []validation.Validation{})

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()
	reader.On("ReadBody", mock.MatchedBy(func(obj *model2.AuthenticatedUser) bool {
		return true
	})).Run(func(args mock.Arguments) {
		args.Get(0).(*model2.AuthenticatedUser).User = model2.User{
			Email:   "test@email.com",
			Name:    "test name",
			Surname: "test surname",
		}
		args.Get(0).(*model2.AuthenticatedUser).Password = "test password"
	}).Return(nil).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.update_user.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}
