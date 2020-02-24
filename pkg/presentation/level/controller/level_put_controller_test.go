package controller_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	levelservice "github.com/arpb2/C-3PO/pkg/domain/service/level"
	"github.com/arpb2/C-3PO/test/mock/golden"
	servicemock "github.com/arpb2/C-3PO/test/mock/service"
	"github.com/stretchr/testify/mock"
	httpcodes "net/http"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	level "github.com/arpb2/C-3PO/pkg/presentation/level/controller"
	httpmock "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/stretchr/testify/assert"
)

func createPutController(service levelservice.Service) controller.Controller {
	return level.CreatePutController(func(ctx *http.Context) {
		// Nothing
	}, service)
}

func TestPutController_IsPut(t *testing.T) {
	assert.Equal(t, "PUT", createPutController(nil).Method)
}

func TestPutControllerPath_IsLevels(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("/levels/:%s", controller.ParamLevelId), createPutController(nil).Path)
}

func TestPutController_GivenNoId_WhenCalled_Then400(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetParameter", controller.ParamLevelId).Return("").Once()

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
	reader.On("GetParameter", controller.ParamLevelId).Return("not uint").Once()

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
	reader.On("GetParameter", controller.ParamLevelId).Return("1000").Once()
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
	reader.On("GetParameter", controller.ParamLevelId).Return("1000").Once()
	reader.On("ReadBody", mock.Anything).Return(nil).Once()

	service := new(servicemock.MockLevelService)
	service.On("StoreLevel", mock.Anything).Return(model.Level{}, expectedErr)

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
	expectedLevel := model.Level{
		Id:          1000,
		Name:        "Some name",
		Description: "Some description",
	}

	reader := new(httpmock.MockReader)
	reader.On("GetParameter", controller.ParamLevelId).Return("1000").Once()
	reader.On("ReadBody", mock.Anything).Run(func(args mock.Arguments) {
		lvl := args.Get(0).(*model.Level)
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
