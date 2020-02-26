package controller_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/user"

	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"
	"github.com/arpb2/C-3PO/test/mock/token"

	usercontroller "github.com/arpb2/C-3PO/pkg/presentation/user/controller"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/auth/user/single"
	"github.com/arpb2/C-3PO/test/mock/golden"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	"github.com/stretchr/testify/assert"
)

func createDeleteController() controller.Controller {
	return usercontroller.CreateDeleteController(
		pipeline2.CreateDebugHttpPipeline(),
		single.CreateMiddleware(
			&token.MockTokenHandler{},
		),
		nil,
	)
}

func TestUserDeleteControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "DELETE", createDeleteController().Method)
}

func TestUserDeleteControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("/users/:%s", user.ParamUserId), createDeleteController().Path)
}

func TestUserDeleteControllerBody_400OnNoUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createDeleteController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserDeleteControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("not a number").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createDeleteController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserDeleteControllerBody_500OnRepositoryDeleteError(t *testing.T) {
	repository := new(repository.MockUserRepository)
	repository.On("DeleteUser", uint(1000)).Return(errors.New("whoops error")).Once()

	body := usercontroller.CreateDeleteBody(pipeline2.CreateDebugHttpPipeline(), repository)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_delete.repository.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestUserDeleteControllerBody_200OnUserDeletedOnRepository(t *testing.T) {
	repository := new(repository.MockUserRepository)
	repository.On("DeleteUser", uint(1000)).Return(nil).Once()

	body := usercontroller.CreateDeleteBody(pipeline2.CreateDebugHttpPipeline(), repository)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	body(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Zero(t, w.Body.Len())
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}
