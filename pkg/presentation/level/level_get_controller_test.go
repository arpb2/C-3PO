package level_test

import (
	"bytes"
	"errors"
	httpcodes "net/http"
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/repository/level"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	level3 "github.com/arpb2/C-3PO/pkg/domain/model/level"
	level2 "github.com/arpb2/C-3PO/pkg/presentation/level"

	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"

	"github.com/arpb2/C-3PO/test/mock/golden"
	repositorymock "github.com/arpb2/C-3PO/test/mock/repository"

	httpmock "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/stretchr/testify/assert"
)

func createGetHandler(repository level.Repository) http.Handler {
	return level2.CreateGetHandler("level_id", pipeline2.CreateDebugHttpPipeline(), repository)
}

func TestGetController_GivenNoId_WhenCalled_Then400(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetParameter", "level_id").Return("").Once()

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createGetHandler(nil)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.empty.level_id.golden.json")

	assert.Equal(t, httpcodes.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestGetController_GivenNoUintId_WhenCalled_Then400(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetParameter", "level_id").Return("not uint").Once()

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createGetHandler(nil)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "bad_request.malformed.level_id.golden.json")

	assert.Equal(t, httpcodes.StatusBadRequest, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
}

func TestGetController_GivenAnErroredRepository_WhenCalled_ThenRepositoryError(t *testing.T) {
	expectedErr := errors.New("error")

	reader := new(httpmock.MockReader)
	reader.On("GetParameter", "level_id").Return("1000").Once()

	repository := new(repositorymock.MockLevelRepository)
	repository.On("GetLevel", uint(1000)).Return(level3.Level{}, expectedErr)

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createGetHandler(repository)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "internal_server_error.error_read.repository.golden.json")

	assert.Equal(t, httpcodes.StatusInternalServerError, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}

func TestGetController_GivenARepositoryWithTheCalleeId_WhenCalled_ThenStoredLevelIsReturned(t *testing.T) {
	expectedLevel := level3.Level{
		Id:          1000,
		Name:        "Some name",
		Description: "Some description",
	}

	reader := new(httpmock.MockReader)
	reader.On("GetParameter", "level_id").Return("1000").Once()

	repository := new(repositorymock.MockLevelRepository)
	repository.On("GetLevel", expectedLevel.Id).Return(expectedLevel, nil)

	c, w := httpmock.CreateTestContext()
	c.Reader = reader

	createGetHandler(repository)(c)

	actual := bytes.TrimSpace([]byte(w.Body.String()))
	expected := golden.Get(t, actual, "ok.get_level.golden.json")

	assert.Equal(t, httpcodes.StatusOK, w.Code)
	assert.Equal(t, expected, actual)
	reader.AssertExpectations(t)
	repository.AssertExpectations(t)
}
