package level_test

import (
	"bytes"
	"errors"
	httpcodes "net/http"
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/repository/level"
	level3 "github.com/arpb2/C-3PO/pkg/domain/model/level"
	level2 "github.com/arpb2/C-3PO/pkg/presentation/level"

	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"

	"github.com/arpb2/C-3PO/test/mock/golden"
	repositorymock "github.com/arpb2/C-3PO/test/mock/repository"
	"github.com/stretchr/testify/mock"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	httpmock "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/stretchr/testify/assert"
)

func createPutHandler(repository level.Repository) http.Handler {
	return level2.CreatePutHandler(
		"level_id",
		pipeline2.CreateDebugHttpPipeline(),
		repository,
	)
}

func TestPutController_GivenNoId_WhenCalled_Then400(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetParameter", "level_id").Return("").Once()
	reader.On("ReadBody", mock.Anything).Return(nil).Maybe()

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createPutHandler(nil)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.level_id.golden.json")

	assert.Equal(t, httpcodes.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestPutController_GivenNoUintId_WhenCalled_Then400(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetParameter", "level_id").Return("not uint").Once()
	reader.On("ReadBody", mock.Anything).Return(nil).Maybe()

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createPutHandler(nil)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.level_id.golden.json")

	assert.Equal(t, httpcodes.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestPutController_GivenNoBody_WhenCalled_Then400(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetParameter", "level_id").Return("1000").Maybe()
	reader.On("ReadBody", mock.Anything).Return(errors.New("error")).Once()

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createPutHandler(nil)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.body.golden.json")

	assert.Equal(t, httpcodes.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestPutController_GivenAnErroredRepository_WhenCalled_ThenRepositoryError(t *testing.T) {
	expectedErr := errors.New("error")

	reader := new(httpmock.MockReader)
	reader.On("GetParameter", "level_id").Return("1000").Once()
	reader.On("ReadBody", mock.Anything).Return(nil).Once()

	repository := new(repositorymock.MockLevelRepository)
	repository.On("StoreLevel", mock.Anything).Return(level3.Level{}, expectedErr)

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createPutHandler(repository)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_store.repository.golden.json")

	assert.Equal(t, httpcodes.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestPutController_GivenARepositoryWithTheCalleeId_WhenCalled_ThenStoredLevelIsReturned(t *testing.T) {
	expectedLevel := level3.Level{
		Id:          1000,
		Name:        "Some name",
		Description: "Some description",
	}

	reader := new(httpmock.MockReader)
	reader.On("GetParameter", "level_id").Return("1000").Once()
	reader.On("ReadBody", mock.Anything).Run(func(args mock.Arguments) {
		lvl := args.Get(0).(*level3.Level)
		lvl.Id = expectedLevel.Id
		lvl.Name = expectedLevel.Name
		lvl.Description = expectedLevel.Description
	}).Return(nil).Once()

	repository := new(repositorymock.MockLevelRepository)
	repository.On("StoreLevel", expectedLevel).Return(expectedLevel, nil)

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createPutHandler(repository)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.put_level.golden.json")

	assert.Equal(t, httpcodes.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}
