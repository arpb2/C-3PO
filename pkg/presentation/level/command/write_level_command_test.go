package command_test

import (
	"errors"
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/level/model"

	levelcommand "github.com/arpb2/C-3PO/pkg/presentation/level/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestWriteUserLevelCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := levelcommand.CreateWriteLevelCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "write_level_command", name)
}

func TestWriteLevelCommand_GivenOneAndAContextWithoutLevel_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(levelcommand.TagLevelId, uint(1000))
	cmd := levelcommand.CreateWriteLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteLevelCommand_GivenOneAndAContextWithoutLevelId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(levelcommand.TagLevel, model2.Level{})
	cmd := levelcommand.CreateWriteLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteLevelCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(levelcommand.TagLevel, model2.Level{})
	ctx.Set(levelcommand.TagLevelId, uint(1000))
	expectedErr := errors.New("some error")
	expectedLevel := model2.Level{}
	s := new(service.MockLevelService)
	s.On("StoreLevel", model2.Level{
		Id: 1000,
	}).Return(expectedLevel, expectedErr)
	cmd := levelcommand.CreateWriteLevelCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestWriteLevelCommand_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(levelcommand.TagLevel, model2.Level{})
	ctx.Set(levelcommand.TagLevelId, uint(1000))
	expectedVal := model2.Level{}
	s := new(service.MockLevelService)
	s.On("StoreLevel", model2.Level{
		Id: 1000,
	}).Return(expectedVal, nil)
	cmd := levelcommand.CreateWriteLevelCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(levelcommand.TagLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
