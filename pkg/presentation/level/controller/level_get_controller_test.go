package controller_test

import (
	"bytes"
	"errors"
	"fmt"
	httpcodes "net/http"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/model"
	levelservice "github.com/arpb2/C-3PO/pkg/domain/service/level"
	"github.com/arpb2/C-3PO/pkg/infra/executor"
	"github.com/arpb2/C-3PO/pkg/infra/pipeline"
	"github.com/arpb2/C-3PO/test/mock/golden"
	servicemock "github.com/arpb2/C-3PO/test/mock/service"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	level "github.com/arpb2/C-3PO/pkg/presentation/level/controller"
	httpmock "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/stretchr/testify/assert"
)

func createGetController(service levelservice.Service) controller.Controller {
	return level.CreateGetController(pipeline.CreateHttpPipeline(executor.CreateDebugHttpExecutor()), service)
}

func TestGetController_IsGet(t *testing.T) {
	assert.Equal(t, "GET", createGetController(nil).Method)
}

func TestGetControllerPath_IsLevels(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("/levels/:%s", controller.ParamLevelId), createGetController(nil).Path)
}

func TestGetController_GivenNoId_WhenCalled_Then400(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetParameter", controller.ParamLevelId).Return("").Once()

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
	reader.On("GetParameter", controller.ParamLevelId).Return("not uint").Once()

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createGetController(nil).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.level_id.golden.json")

	assert.Equal(t, httpcodes.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestGetController_GivenAnErroredService_WhenCalled_ThenServiceError(t *testing.T) {
	expectedErr := errors.New("error")

	reader := new(httpmock.MockReader)
	reader.On("GetParameter", controller.ParamLevelId).Return("1000").Once()

	service := new(servicemock.MockLevelService)
	service.On("GetLevel", uint(1000)).Return(model.Level{}, expectedErr)

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createGetController(service).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_read.service.golden.json")

	assert.Equal(t, httpcodes.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}

func TestGetController_GivenAServiceWithTheCalleeId_WhenCalled_ThenStoredLevelIsReturned(t *testing.T) {
	expectedLevel := model.Level{
		Id:          1000,
		Name:        "Some name",
		Description: "Some description",
	}

	reader := new(httpmock.MockReader)
	reader.On("GetParameter", controller.ParamLevelId).Return("1000").Once()

	service := new(servicemock.MockLevelService)
	service.On("GetLevel", expectedLevel.Id).Return(expectedLevel, nil)

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createGetController(service).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_level.golden.json")

	assert.Equal(t, httpcodes.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}
