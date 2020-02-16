package user_level_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"testing"

	controller2 "github.com/arpb2/C-3PO/pkg/controller"

	"github.com/arpb2/C-3PO/pkg/pipeline"

	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/auth/jwt"
	userlevelcontroller "github.com/arpb2/C-3PO/pkg/controller/user_level"
	"github.com/arpb2/C-3PO/pkg/executor"
	"github.com/arpb2/C-3PO/pkg/middleware/auth/teacher"
	teacherservice "github.com/arpb2/C-3PO/pkg/service/teacher"
	userservice "github.com/arpb2/C-3PO/pkg/service/user"
	userlevelservice "github.com/arpb2/C-3PO/pkg/service/user_level"
	"github.com/arpb2/C-3PO/test/mock/golden"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/service"
	"github.com/stretchr/testify/assert"
)

func createGetController() controller.Controller {
	return userlevelcontroller.CreateGetController(
		pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()),
		teacher.CreateMiddleware(
			jwt.CreateTokenHandler(),
			teacherservice.CreateService(userservice.CreateService()),
		),
		userlevelservice.CreateService(),
	)
}

func TestCodeGetControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", createGetController().Method)
}

func TestCodeGetControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("/users/:%s/levels/:%s", controller2.ParamUserId, controller2.ParamLevelId), createGetController().Path)
}

func TestCodeGetControllerBody_400OnEmptyUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller2.ParamLevelId).Return("1").Maybe()
	reader.On("GetParameter", controller2.ParamUserId).Return("").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestCodeGetControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller2.ParamLevelId).Return("1000").Maybe()
	reader.On("GetParameter", controller2.ParamUserId).Return("not a number").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestCodeGetControllerBody_400OnMalformedLevelId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller2.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", controller2.ParamLevelId).Return("not a number").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.level_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestCodeGetControllerBody_400OnEmptyLevelId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller2.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", controller2.ParamLevelId).Return("").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	createGetController().Body(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.level_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestCodeGetControllerBody_500OnServiceReadError(t *testing.T) {
	userLevelService := new(service.MockUserLevelService)
	userLevelService.On("GetUserLevel", uint(1000), uint(1000)).Return(nil, errors.New("whoops error"))

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller2.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", controller2.ParamLevelId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	userlevelcontroller.CreateGetBody(pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()), userLevelService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_read.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	userLevelService.AssertExpectations(t)
}

func TestCodeGetControllerBody_400OnNoCodeStoredInService(t *testing.T) {
	userLevelService := new(service.MockUserLevelService)
	userLevelService.On("GetUserLevel", uint(1000), uint(1000)).Return(nil, nil)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller2.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", controller2.ParamLevelId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	userlevelcontroller.CreateGetBody(pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()), userLevelService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "not_found.missing_user_level.read.service.golden.json")

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	userLevelService.AssertExpectations(t)
}

func TestCodeGetControllerBody_200OnCodeStoredOnService(t *testing.T) {
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
	userLevelService.On("GetUserLevel", uint(1000), uint(1000)).Return(&model.UserLevel{
		UserId:  1000,
		LevelId: 1000,
		Code:    expectedCode,
	}, nil)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller2.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", controller2.ParamLevelId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	userlevelcontroller.CreateGetBody(pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()), userLevelService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_user_level.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	userLevelService.AssertExpectations(t)
}

func TestCodeGetControllerBody_200OnEmptyCodeStoredOnService(t *testing.T) {
	expectedCode := ""
	userLevelService := new(service.MockUserLevelService)
	userLevelService.On("GetUserLevel", uint(1000), uint(1000)).Return(&model.UserLevel{
		UserId:  1000,
		LevelId: 1000,
		Code:    expectedCode,
	}, nil)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller2.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", controller2.ParamLevelId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	userlevelcontroller.CreateGetBody(pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()), userLevelService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_empty_user_level.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	userLevelService.AssertExpectations(t)
}