package command_test

import (
	"errors"
	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	command2 "github.com/arpb2/C-3PO/pkg/presentation/user_level/command"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestWriteUserLevelCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command2.CreateWriteUserLevelCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "write_user_level_command", name)
}

func TestWriteUserLevelCommand_GivenOneAndAContextWithoutRawCode_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, uint(1000))
	ctx.Set(command2.TagLevelId, uint(1000))
	cmd := command2.CreateWriteUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteUserLevelCommand_GivenOneAndAContextWithoutLevelId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, uint(1000))
	ctx.Set(command2.TagUserLevelData, model.UserLevelData{
		Code: "code",
	})
	cmd := command2.CreateWriteUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteUserLevelCommand_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command2.TagUserLevelData, model.UserLevelData{})
	ctx.Set(command2.TagLevelId, uint(1000))
	cmd := command2.CreateWriteUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteUserLevelCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command2.TagUserLevelData, model.UserLevelData{
		Code: "code",
	})
	ctx.Set(command2.TagLevelId, uint(1000))
	ctx.Set(command.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	expectedUserLevel := model.UserLevel{
		UserLevelData: model.UserLevelData{
			Code: "code",
		},
		LevelId: uint(1000),
		UserId:  uint(1000),
	}
	s := new(service.MockUserLevelService)
	s.On("StoreUserLevel", expectedUserLevel).Return(expectedUserLevel, expectedErr)
	cmd := command2.CreateWriteUserLevelCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestWriteUserLevelCommand_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command2.TagUserLevelData, model.UserLevelData{
		Code: "code",
	})
	ctx.Set(command2.TagLevelId, uint(1000))
	ctx.Set(command.TagUserId, uint(1000))
	expectedVal := model.UserLevel{
		UserLevelData: model.UserLevelData{
			Code: "code",
		},
		LevelId: uint(1000),
		UserId:  uint(1000),
	}
	s := new(service.MockUserLevelService)
	s.On("StoreUserLevel", expectedVal).Return(expectedVal, nil)
	cmd := command2.CreateWriteUserLevelCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command2.TagUserLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
