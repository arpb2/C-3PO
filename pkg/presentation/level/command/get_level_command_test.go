package command_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/level/command"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestGetLevelCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateGetLevelCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "get_level_command", name)
}

func TestGetLevelCommand_GivenOneAndAContextWithoutAuthenticatedLevel_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateGetLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestGetLevelCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	expectedVal := uint(1000)
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagLevelId, expectedVal)
	expectedErr := errors.New("some error")
	s := new(service.MockLevelService)
	s.On("GetLevel", expectedVal).Return(model.Level{}, expectedErr)
	cmd := command.CreateGetLevelCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetLevelCommand_GivenOneAndANoLevelCreatedService_WhenRunning_Then404(t *testing.T) {
	expectedVal := uint(1000)
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagLevelId, expectedVal)
	expectedErr := http.CreateNotFoundError()
	s := new(service.MockLevelService)
	s.On("GetLevel", expectedVal).Return(model.Level{}, http.CreateNotFoundError())
	cmd := command.CreateGetLevelCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetLevelCommand_GivenOne_WhenRunning_ThenContextHasLevelAndReturnsNoError(t *testing.T) {
	expectedVal := model.Level{}
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagLevelId, uint(1000))
	s := new(service.MockLevelService)
	s.On("GetLevel", uint(1000)).Return(expectedVal, nil)
	cmd := command.CreateGetLevelCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command.TagLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
