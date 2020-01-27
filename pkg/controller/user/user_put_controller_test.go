package user_controller_test

import (
	"bytes"
	"errors"
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/auth/jwt"
	"github.com/arpb2/C-3PO/pkg/controller/user"
	"github.com/arpb2/C-3PO/pkg/executor"
	"github.com/arpb2/C-3PO/pkg/middleware/auth/single_auth"
	user_service "github.com/arpb2/C-3PO/pkg/service/user"
	user_validation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/arpb2/C-3PO/hack/golden"
	test_http_wrapper "github.com/arpb2/C-3PO/hack/http_wrapper"
	"github.com/arpb2/C-3PO/hack/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func createPutController() controller.Controller {
	return user_controller.CreatePutController(
		executor.CreatePipeline(executor.CreateDebugHttpExecutor()),
		[]user_validation.Validation{},
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
	reader := new(test_http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil).Maybe()
	reader.On("GetParameter", "user_id").Return("").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserPutControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(test_http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil).Maybe()
	reader.On("GetParameter", "user_id").Return("not a number").Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserPutControllerBody_400OnEmptyOrMalformedUser(t *testing.T) {
	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(errors.New("malformed")).Once()

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserPutControllerBody_500OnServiceCreateError(t *testing.T) {
	service := new(service.MockUserService)
	service.On("UpdateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(&model.User{
		Id:      1234,
		Email:   "test@email.com",
		Name:    "test name",
		Surname: "test surname",
	}, errors.New("whoops error")).Once()

	body := user_controller.CreatePutBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), []user_validation.Validation{}, service)

	reader := new(test_http_wrapper.MockReader)
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

	c, w := test_http_wrapper.CreateTestContext()
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
	service := new(service.MockUserService)
	service.On("UpdateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil, nil).Once()

	body := user_controller.CreatePutBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), []user_validation.Validation{}, service)

	reader := new(test_http_wrapper.MockReader)
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

	c, w := test_http_wrapper.CreateTestContext()
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
	service := new(service.MockUserService)

	body := user_controller.CreatePutBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), []user_validation.Validation{
		user_validation.IdProvided,
	}, service)

	reader := new(test_http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		obj.User = &model.User{
			Id: 1000,
		}
		return true
	})).Return(nil).Once()

	c, w := test_http_wrapper.CreateTestContext()
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
	service := new(service.MockUserService)
	service.On("UpdateUser", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(expectedUser, nil).Once()

	body := user_controller.CreatePutBody(executor.CreatePipeline(executor.CreateDebugHttpExecutor()), []user_validation.Validation{}, service)

	reader := new(test_http_wrapper.MockReader)
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

	c, w := test_http_wrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.update_user.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}
