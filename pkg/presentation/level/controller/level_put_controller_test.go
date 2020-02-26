package controller_test

import (
	"bytes"
	"errors"
	"fmt"
	level2 "github.com/arpb2/C-3PO/pkg/presentation/level"
	httpcodes "net/http"
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/level/model"
	"github.com/arpb2/C-3PO/pkg/domain/level/service"
	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"

	"github.com/arpb2/C-3PO/test/mock/golden"
	servicemock "github.com/arpb2/C-3PO/test/mock/service"
	"github.com/stretchr/testify/mock"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	level "github.com/arpb2/C-3PO/pkg/presentation/level/controller"
	httpmock "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/stretchr/testify/assert"
)

func createPutController(service service.Service) controller.Controller {
	return level.CreatePutController(
		pipeline2.CreateDebugHttpPipeline(),
		func(ctx *http.Context) {
			// Nothing
		},
		service,
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

func TestPutController_GivenAnErroredService_WhenCalled_ThenServiceError(t *testing.T) {
	expectedErr := errors.New("error")

	reader := new(httpmock.MockReader)
	reader.On("GetParameter", level2.ParamLevelId).Return("1000").Once()
	reader.On("ReadBody", mock.Anything).Return(nil).Once()

	service := new(servicemock.MockLevelService)
	service.On("StoreLevel", mock.Anything).Return(model2.Level{}, expectedErr)

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createPutController(service).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_store.service.golden.json")

	assert.Equal(t, httpcodes.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}

func TestPutController_GivenAServiceWithTheCalleeId_WhenCalled_ThenStoredLevelIsReturned(t *testing.T) {
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

	service := new(servicemock.MockLevelService)
	service.On("StoreLevel", expectedLevel).Return(expectedLevel, nil)

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createPutController(service).Body(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.put_level.golden.json")

	assert.Equal(t, httpcodes.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	service.AssertExpectations(t)
}
