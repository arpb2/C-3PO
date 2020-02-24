package command_test

import (
	"errors"
	"testing"

	levelcommand "github.com/arpb2/C-3PO/pkg/presentation/level/command"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	usercommand "github.com/arpb2/C-3PO/pkg/presentation/user_level/command"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestGetUserLevelCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := usercommand.CreateGetUserLevelCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "get_user_level_command", name)
}

func TestGetUserLevelCommand_GivenOneAndAContextWithoutLevelID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, uint(1000))
	cmd := usercommand.CreateGetUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestGetUserLevelCommand_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(levelcommand.TagLevelId, uint(1000))
	cmd := usercommand.CreateGetUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestGetUserLevelCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(levelcommand.TagLevelId, uint(1000))
	ctx.Set(command.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	s := new(service.MockUserLevelService)
	s.On("GetUserLevel", uint(1000), uint(1000)).Return(model.UserLevel{}, expectedErr)
	cmd := usercommand.CreateGetUserLevelCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetUserLevelCommand_GivenOneAndAServiceWithNoCode_WhenRunning_Then404(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(levelcommand.TagLevelId, uint(1000))
	ctx.Set(command.TagUserId, uint(1000))
	s := new(service.MockUserLevelService)
	s.On("GetUserLevel", uint(1000), uint(1000)).Return(model.UserLevel{}, http.CreateNotFoundError())
	cmd := usercommand.CreateGetUserLevelCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateNotFoundError(), err)
	s.AssertExpectations(t)
}

func TestGetUserLevelCommand_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(levelcommand.TagLevelId, uint(1000))
	ctx.Set(command.TagUserId, uint(1000))
	expectedVal := model.UserLevel{}
	s := new(service.MockUserLevelService)
	s.On("GetUserLevel", uint(1000), uint(1000)).Return(expectedVal, nil)
	cmd := usercommand.CreateGetUserLevelCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(usercommand.TagUserLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
