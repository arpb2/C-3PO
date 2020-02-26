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

	http2 "github.com/arpb2/C-3PO/pkg/domain/architecture/http"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/auth/user/single"
	"github.com/arpb2/C-3PO/test/mock/golden"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	"github.com/stretchr/testify/assert"
)

func createGetController() controller.Controller {
	return usercontroller.CreateGetController(
		pipeline2.CreateDebugHttpPipeline(),
		single.CreateMiddleware(
			&token.MockTokenHandler{},
		),
		nil,
	)
}

func TestUserGetControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", createGetController().Method)
}

func TestUserGetControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("/users/:%s", user.ParamUserId), createGetController().Path)
}

func TestUserGetControllerBody_400OnEmptyUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("").Once()

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
	reader.On("GetParameter", user.ParamUserId).Return("not a number").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserGetControllerBody_500OnRepositoryReadError(t *testing.T) {
	repository := new(repository.MockUserRepository)
	repository.On("GetUser", uint(1000)).Return(nil, errors.New("whoops error")).Once()

	body := usercontroller.CreateGetBody(pipeline2.CreateDebugHttpPipeline(), repository)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_read.repository.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestUserGetControllerBody_400OnNoUserStoredInRepository(t *testing.T) {
	repository := new(repository.MockUserRepository)
	repository.On("GetUser", uint(1000)).Return(model2.User{}, http2.CreateNotFoundError()).Once()

	body := usercontroller.CreateGetBody(pipeline2.CreateDebugHttpPipeline(), repository)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "not_found.missing_user.read.repository.golden.json")

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestUserGetControllerBody_200OnUserStoredOnRepository(t *testing.T) {
	expectedUser := model2.User{
		Id:      1000,
		Email:   "test@email.com",
		Name:    "TestName",
		Surname: "TestSurname",
	}
	repository := new(repository.MockUserRepository)
	repository.On("GetUser", uint(1000)).Return(expectedUser, nil).Once()

	body := usercontroller.CreateGetBody(pipeline2.CreateDebugHttpPipeline(), repository)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_user.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}
