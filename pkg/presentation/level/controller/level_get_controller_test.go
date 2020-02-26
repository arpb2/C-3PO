package controller_test

import (
	"bytes"
	"errors"
	"fmt"
	httpcodes "net/http"
	"testing"

	level2 "github.com/arpb2/C-3PO/pkg/presentation/level"

	model2 "github.com/arpb2/C-3PO/pkg/domain/level/model"
	"github.com/arpb2/C-3PO/pkg/domain/level/repository"
	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"

	"github.com/arpb2/C-3PO/test/mock/golden"
	repositorymock "github.com/arpb2/C-3PO/test/mock/repository"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	level "github.com/arpb2/C-3PO/pkg/presentation/level/controller"
	httpmock "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/stretchr/testify/assert"
)

func createGetController(repository repository.LevelRepository) controller.Controller {
	return level.CreateGetController(pipeline2.CreateDebugHttpPipeline(), repository)
}

func TestGetController_IsGet(t *testing.T) {
	assert.Equal(t, "GET", createGetController(nil).Method)
}

func TestGetControllerPath_IsLevels(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("/levels/:%s", level2.ParamLevelId), createGetController(nil).Path)
}

func TestGetController_GivenNoId_WhenCalled_Then400(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetParameter", level2.ParamLevelId).Return("").Once()

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createGetController(nil).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.level_id.golden.json")

	assert.Equal(t, httpcodes.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestGetController_GivenNoUintId_WhenCalled_Then400(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetParameter", level2.ParamLevelId).Return("not uint").Once()

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createGetController(nil).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.level_id.golden.json")

	assert.Equal(t, httpcodes.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestGetController_GivenAnErroredRepository_WhenCalled_ThenRepositoryError(t *testing.T) {
	expectedErr := errors.New("error")

	reader := new(httpmock.MockReader)
	reader.On("GetParameter", level2.ParamLevelId).Return("1000").Once()

	repository := new(repositorymock.MockLevelRepository)
	repository.On("GetLevel", uint(1000)).Return(model2.Level{}, expectedErr)

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createGetController(repository).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_read.repository.golden.json")

	assert.Equal(t, httpcodes.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestGetController_GivenARepositoryWithTheCalleeId_WhenCalled_ThenStoredLevelIsReturned(t *testing.T) {
	expectedLevel := model2.Level{
		Id:          1000,
		Name:        "Some name",
		Description: "Some description",
	}

	reader := new(httpmock.MockReader)
	reader.On("GetParameter", level2.ParamLevelId).Return("1000").Once()

	repository := new(repositorymock.MockLevelRepository)
	repository.On("GetLevel", expectedLevel.Id).Return(expectedLevel, nil)

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createGetController(repository).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_level.golden.json")

	assert.Equal(t, httpcodes.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}
