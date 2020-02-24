package controller_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"testing"

	controller3 "github.com/arpb2/C-3PO/pkg/presentation/user/controller"

	http2 "github.com/arpb2/C-3PO/pkg/domain/http"

	"github.com/arpb2/C-3PO/pkg/infra/pipeline"

	"github.com/arpb2/C-3PO/pkg/data/jwt"
	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/arpb2/C-3PO/pkg/infra/executor"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user/single"
	"github.com/arpb2/C-3PO/test/mock/golden"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/service"
	"github.com/stretchr/testify/assert"
)

func createGetController() controller.Controller {
	return controller3.CreateGetController(
		pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()),
		single.CreateMiddleware(
			jwt.CreateTokenHandler([]byte("52bfd2de0a2e69dff4517518590ac32a46bd76606ec22a258f99584a6e70aca2")),
		),
		nil,
	)
}

func TestUserGetControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", createGetController().Method)
}

func TestUserGetControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("/users/:%s", controller.ParamUserId), createGetController().Path)
}

func TestUserGetControllerBody_400OnEmptyUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller.ParamUserId).Return("").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserGetControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller.ParamUserId).Return("not a number").Once()

	c, w := testhttpwrapper.CreateTestContext()
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

	body := controller3.CreateGetBody(pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()), service)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller.ParamUserId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
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
	service.On("GetUser", uint(1000)).Return(model.User{}, http2.CreateNotFoundError()).Once()

	body := controller3.CreateGetBody(pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()), service)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller.ParamUserId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
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
	expectedUser := model.User{
		Id:      1000,
		Email:   "test@email.com",
		Name:    "TestName",
		Surname: "TestSurname",
	}
	service := new(service.MockUserService)
	service.On("GetUser", uint(1000)).Return(expectedUser, nil).Once()

	body := controller3.CreateGetBody(pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()), service)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller.ParamUserId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_user.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}
