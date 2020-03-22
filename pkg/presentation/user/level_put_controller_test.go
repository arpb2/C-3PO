package user_test

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"

	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"
	"github.com/arpb2/C-3PO/pkg/presentation/user"

	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"

	"github.com/arpb2/C-3PO/test/mock/golden"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	"github.com/stretchr/testify/assert"
)

func TestUserLevelPutControllerBody_400OnEmptyUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.Anything).Return(nil).Once()
	reader.On("GetParameter", "level_id").Return("1000").Once()
	reader.On("GetParameter", "user_id").Return("").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	user.CreatePutLevelHandler("user_id", "level_id", pipeline2.CreateDebugHttpPipeline(), nil)(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserLevelPutControllerBody_400OnMalformedUserId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.Anything).Return(nil).Maybe()
	reader.On("GetParameter", "level_id").Return("1000").Maybe()
	reader.On("GetParameter", "user_id").Return("not a number").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	user.CreatePutLevelHandler("user_id", "level_id", pipeline2.CreateDebugHttpPipeline(), nil)(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.user_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserLevelPutControllerBody_400OnMalformedLevelId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.Anything).Return(nil).Maybe()
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "level_id").Return("not a number").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	user.CreatePutLevelHandler("user_id", "level_id", pipeline2.CreateDebugHttpPipeline(), nil)(c)
	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.level_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserLevelPutControllerBody_400OnEmptyLevelId(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.Anything).Return(nil).Maybe()
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "level_id").Return("").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	user.CreatePutLevelHandler("user_id", "level_id", pipeline2.CreateDebugHttpPipeline(), nil)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.level_id.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserLevelPutControllerBody_400OnNoCode(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "level_id").Return("1000").Once()
	reader.On("ReadBody", mock.Anything).Return(errors.New("no code")).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	user.CreatePutLevelHandler("user_id", "level_id", pipeline2.CreateDebugHttpPipeline(), nil)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.code.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserLevelPutControllerBody_400OnNoWorkspace(t *testing.T) {
	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "level_id").Return("1000").Once()
	reader.On("ReadBody", mock.Anything).Return(errors.New("no workspace")).Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	user.CreatePutLevelHandler("user_id", "level_id", pipeline2.CreateDebugHttpPipeline(), nil)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.workspace.golden.json")

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestUserLevelPutControllerBody_500OnRepositoryWriteError(t *testing.T) {
	userLevelRepository := new(repository.MockUserLevelRepository)
	userLevelRepository.On("GetUserLevel", uint(1000), uint(1000)).Return(user2.Level{}, errors.New("whoops error"))

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "level_id").Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	user.CreateGetLevelHandler("user_id", "level_id", pipeline2.CreateDebugHttpPipeline(), userLevelRepository)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_write.repository.golden.json")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	userLevelRepository.AssertExpectations(t)
	reader.AssertExpectations(t)
}

func TestUserLevelPutControllerBody_200OnUserLevelReplacedOnRepository(t *testing.T) {
	expectedCode := `
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello world!")
}
			`
	userLevelRepository := new(repository.MockUserLevelRepository)
	userLevelRepository.On("GetUserLevel", uint(1000), uint(1000)).Return(user2.Level{
		UserId:  1000,
		LevelId: 1000,
		LevelData: user2.LevelData{
			Code: expectedCode,
		},
	}, nil)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "level_id").Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	user.CreateGetLevelHandler("user_id", "level_id", pipeline2.CreateDebugHttpPipeline(), userLevelRepository)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.replace_user_level.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	userLevelRepository.AssertExpectations(t)
	reader.AssertExpectations(t)
}

func TestUserLevelPutControllerBody_200OnEmptyUserLevelStoredOnRepository(t *testing.T) {
	expectedCode := ""
	userLevelRepository := new(repository.MockUserLevelRepository)
	userLevelRepository.On("GetUserLevel", uint(1000), uint(1000)).Return(user2.Level{
		UserId:  1000,
		LevelId: 1000,
		LevelData: user2.LevelData{
			Code: expectedCode,
		},
	}, nil)

	reader := new(testhttpwrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1000").Once()
	reader.On("GetParameter", "level_id").Return("1000").Once()

	c, w := testhttpwrapper.CreateTestContext()
	c.Reader = reader

	user.CreateGetLevelHandler("user_id", "level_id", pipeline2.CreateDebugHttpPipeline(), userLevelRepository)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.replace_empty_user_level.golden.json")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	userLevelRepository.AssertExpectations(t)
	reader.AssertExpectations(t)
}
