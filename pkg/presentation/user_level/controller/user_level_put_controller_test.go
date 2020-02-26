package controller_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/arpb2/C-3PO/pkg/presentation/level"
	"github.com/arpb2/C-3PO/pkg/presentation/user"
	"net/http"
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user_level/model"
	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"
	"github.com/arpb2/C-3PO/test/mock/token"

	userlevelcontroller "github.com/arpb2/C-3PO/pkg/presentation/user_level/controller"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user/teacher"
	"github.com/arpb2/C-3PO/test/mock/golden"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/service"
	"github.com/stretchr/testify/assert"
)

func createPutController() controller.Controller {
	return userlevelcontroller.CreatePutController(
		pipeline2.CreateDebugHttpPipeline(),
		teacher.CreateMiddleware(
			&token.MockTokenHandler{},
			nil,
		),
		nil,
	)
}

func TestUserLevelPutControllerMethodIsPUT(t *testing.T) {
	assert.Equal(t, "PUT", createPutController().Method)
}

func TestUserLevelPutControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("/users/:%s/levels/:%s", user.ParamUserId, level.ParamLevelId), createPutController().Path)
}

func TestUserLevelPutControllerBody_400OnEmptyUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetFormData", "code").Return("", true).Once()
	reader.On("GetFormData", "workspace").Return("", true).Once()
	reader.On("GetParameter", level.ParamLevelId).Return("1000").Once()
	reader.On("GetParameter", user.ParamUserId).Return("").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
}

func TestUserLevelPutControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetFormData", "code").Return("code", true).Maybe()
	reader.On("GetFormData", "workspace").Return("workspace", true).Maybe()
	reader.On("GetParameter", level.ParamLevelId).Return("1000").Maybe()
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

func TestUserLevelPutControllerBody_400OnMalformedLevelId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetFormData", "code").Return("code", true).Maybe()
	reader.On("GetFormData", "workspace").Return("workspace", true).Maybe()
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", level.ParamLevelId).Return("not a number").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.level_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserLevelPutControllerBody_400OnEmptyLevelId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetFormData", "code").Return("code", true).Maybe()
	reader.On("GetFormData", "workspace").Return("workspace", true).Maybe()
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", level.ParamLevelId).Return("").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.level_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
}

func TestUserLevelPutControllerBody_400OnNoCode(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", level.ParamLevelId).Return("1000").Once()
	reader.On("GetFormData", "code").Return("", false).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.code.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
}

func TestUserLevelPutControllerBody_400OnNoWorkspace(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", level.ParamLevelId).Return("1000").Once()
	reader.On("GetFormData", "code").Return("code", true).Once()
	reader.On("GetFormData", "workspace").Return("", false).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createPutController().Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.workspace.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
}

func TestUserLevelPutControllerBody_500OnServiceWriteError(t *testing.T) {
	userLevelService := new(service.MockUserLevelService)
	userLevelService.On("GetUserLevel", uint(1000), uint(1000)).Return(model2.UserLevel{}, errors.New("whoops error"))

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", level.ParamLevelId).Return("1000").Once()
	reader.On("GetFormData", "code").Return("sending some code", true).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	userlevelcontroller.CreateGetBody(pipeline2.CreateDebugHttpPipeline(), userLevelService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_write.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	userLevelService.AssertExpectations(t)
}

func TestUserLevelPutControllerBody_200OnUserLevelReplacedOnService(t *testing.T) {
	expectedCode := `
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello world!")
}
			`
	userLevelService := new(service.MockUserLevelService)
	userLevelService.On("GetUserLevel", uint(1000), uint(1000)).Return(model2.UserLevel{
		UserId:  1000,
		LevelId: 1000,
		UserLevelData: model2.UserLevelData{
			Code: expectedCode,
		},
	}, nil)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", level.ParamLevelId).Return("1000").Once()
	reader.On("GetFormData", "code").Return(expectedCode, true).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	userlevelcontroller.CreateGetBody(pipeline2.CreateDebugHttpPipeline(), userLevelService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.replace_user_level.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	userLevelService.AssertExpectations(t)
}

func TestUserLevelPutControllerBody_200OnEmptyUserLevelStoredOnService(t *testing.T) {
	expectedCode := ""
	userLevelService := new(service.MockUserLevelService)
	userLevelService.On("GetUserLevel", uint(1000), uint(1000)).Return(model2.UserLevel{
		UserId:  1000,
		LevelId: 1000,
		UserLevelData: model2.UserLevelData{
			Code: expectedCode,
		},
	}, nil)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", level.ParamLevelId).Return("1000").Once()
	reader.On("GetFormData", "code").Return(expectedCode, true).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	userlevelcontroller.CreateGetBody(pipeline2.CreateDebugHttpPipeline(), userLevelService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.replace_empty_user_level.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	userLevelService.AssertExpectations(t)
}
