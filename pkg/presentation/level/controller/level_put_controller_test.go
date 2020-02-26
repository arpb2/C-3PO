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
	"github.com/stretchr/testify/mock"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	level "github.com/arpb2/C-3PO/pkg/presentation/level/controller"
	httpmock "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/stretchr/testify/assert"
)

func createPutController(repository repository.LevelRepository) controller.Controller {
	return level.CreatePutController(
		pipeline2.CreateDebugHttpPipeline(),
		func(ctx *http.Context) {
			// Nothing
		},
		repository,
	)
}

func TestPutController_IsPut(t *testing.T) {
	assert.Equal(t, "PUT", createPutController(nil).Method)
}

func TestPutControllerPath_IsLevels(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("/levels/:%s", level2.ParamLevelId), createPutController(nil).Path)
}

func TestPutController_GivenNoId_WhenCalled_Then400(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetParameter", level2.ParamLevelId).Return("").Once()
	reader.On("ReadBody", mock.Anything).Return(nil).Maybe()

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createPutController(nil).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.level_id.golden.json")

	assert.Equal(t, httpcodes.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestPutController_GivenNoUintId_WhenCalled_Then400(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetParameter", level2.ParamLevelId).Return("not uint").Once()
	reader.On("ReadBody", mock.Anything).Return(nil).Maybe()

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createPutController(nil).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.level_id.golden.json")

	assert.Equal(t, httpcodes.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestPutController_GivenNoBody_WhenCalled_Then400(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetParameter", level2.ParamLevelId).Return("1000").Maybe()
	reader.On("ReadBody", mock.Anything).Return(errors.New("error")).Once()

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createPutController(nil).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.body.golden.json")

	assert.Equal(t, httpcodes.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestPutController_GivenAnErroredRepository_WhenCalled_ThenRepositoryError(t *testing.T) {
	expectedErr := errors.New("error")

	reader := new(httpmock.MockReader)
	reader.On("GetParameter", level2.ParamLevelId).Return("1000").Once()
	reader.On("ReadBody", mock.Anything).Return(nil).Once()

	repository := new(repositorymock.MockLevelRepository)
	repository.On("StoreLevel", mock.Anything).Return(model2.Level{}, expectedErr)

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createPutController(repository).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_store.repository.golden.json")

	assert.Equal(t, httpcodes.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestPutController_GivenARepositoryWithTheCalleeId_WhenCalled_ThenStoredLevelIsReturned(t *testing.T) {
	expectedLevel := model2.Level{
		Id:          1000,
		Name:        "Some name",
		Description: "Some description",
	}

	reader := new(httpmock.MockReader)
	reader.On("GetParameter", level2.ParamLevelId).Return("1000").Once()
	reader.On("ReadBody", mock.Anything).Run(func(args mock.Arguments) {
		lvl := args.Get(0).(*model2.Level)
		lvl.Id = expectedLevel.Id
		lvl.Name = expectedLevel.Name
		lvl.Description = expectedLevel.Description
	}).Return(nil).Once()

	repository := new(repositorymock.MockLevelRepository)
	repository.On("StoreLevel", expectedLevel).Return(expectedLevel, nil)

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createPutController(repository).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.put_level.golden.json")

	assert.Equal(t, httpcodes.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}
