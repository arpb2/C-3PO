package controller_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"testing"

	controller3 "github.com/arpb2/C-3PO/pkg/domain/level/controller"
	controller2 "github.com/arpb2/C-3PO/pkg/domain/user/controller"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user_level/model"
	mockpipeline "github.com/arpb2/C-3PO/test/mock/pipeline"
	"github.com/arpb2/C-3PO/test/mock/token"

	userlevelcontroller "github.com/arpb2/C-3PO/pkg/presentation/user_level/controller"

	http2 "github.com/arpb2/C-3PO/pkg/domain/architecture/http"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user/teacher"
	"github.com/arpb2/C-3PO/test/mock/golden"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/service"
	"github.com/stretchr/testify/assert"
)

func createGetController() controller.Controller {
	return userlevelcontroller.CreateGetController(
		mockpipeline.CreateDebugHttpPipeline(),
		teacher.CreateMiddleware(
			&token.MockTokenHandler{},
			nil,
		),
		nil,
	)
}

func TestCodeGetControllerMethodIsGET(t *testing.T) {
	assert.Equal(t, "GET", createGetController().Method)
}

func TestCodeGetControllerPathIsAsExpected(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("/users/:%s/levels/:%s", controller2.ParamUserId, controller3.ParamLevelId), createGetController().Path)
}

func TestCodeGetControllerBody_400OnEmptyUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller3.ParamLevelId).Return("1").Maybe()
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
	reader.On("GetParameter", controller3.ParamLevelId).Return("1000").Maybe()
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
	reader.On("GetParameter", controller3.ParamLevelId).Return("not a number").Once()

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
	reader.On("GetParameter", controller3.ParamLevelId).Return("").Once()

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
	userLevelService.On("GetUserLevel", uint(1000), uint(1000)).Return(model2.UserLevel{}, errors.New("whoops error"))

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller2.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", controller3.ParamLevelId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	userlevelcontroller.CreateGetBody(mockpipeline.CreateDebugHttpPipeline(), userLevelService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_read.service.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	userLevelService.AssertExpectations(t)
}

func TestCodeGetControllerBody_400OnNoCodeStoredInService(t *testing.T) {
	userLevelService := new(service.MockUserLevelService)
	userLevelService.On("GetUserLevel", uint(1000), uint(1000)).Return(model2.UserLevel{}, http2.CreateNotFoundError())

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller2.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", controller3.ParamLevelId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	userlevelcontroller.CreateGetBody(mockpipeline.CreateDebugHttpPipeline(), userLevelService)(c)

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
	userLevelService.On("GetUserLevel", uint(1000), uint(1000)).Return(model2.UserLevel{
		UserId:  1000,
		LevelId: 1000,
		UserLevelData: model2.UserLevelData{
			Code: expectedCode,
		},
	}, nil)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller2.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", controller3.ParamLevelId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	userlevelcontroller.CreateGetBody(mockpipeline.CreateDebugHttpPipeline(), userLevelService)(c)

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
	userLevelService.On("GetUserLevel", uint(1000), uint(1000)).Return(model2.UserLevel{
		UserId:  1000,
		LevelId: 1000,
		UserLevelData: model2.UserLevelData{
			Code: expectedCode,
		},
	}, nil)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", controller2.ParamUserId).Return("1000").Once()
	reader.On("GetParameter", controller3.ParamLevelId).Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	userlevelcontroller.CreateGetBody(mockpipeline.CreateDebugHttpPipeline(), userLevelService)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_empty_user_level.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	userLevelService.AssertExpectations(t)
}
