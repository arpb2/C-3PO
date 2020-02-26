package command_test

import (
	"errors"
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user_level/model"

	levelcommand "github.com/arpb2/C-3PO/pkg/presentation/level/command"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	userlevelcommand "github.com/arpb2/C-3PO/pkg/presentation/user_level/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestWriteUserLevelCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := userlevelcommand.CreateWriteUserLevelCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "write_user_level_command", name)
}

func TestWriteUserLevelCommand_GivenOneAndAContextWithoutRawCode_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, uint(1000))
	ctx.Set(levelcommand.TagLevelId, uint(1000))
	cmd := userlevelcommand.CreateWriteUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteUserLevelCommand_GivenOneAndAContextWithoutLevelId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, uint(1000))
	ctx.Set(userlevelcommand.TagUserLevelData, model2.UserLevelData{
		Code: "code",
	})
	cmd := userlevelcommand.CreateWriteUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteUserLevelCommand_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(userlevelcommand.TagUserLevelData, model2.UserLevelData{})
	ctx.Set(levelcommand.TagLevelId, uint(1000))
	cmd := userlevelcommand.CreateWriteUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteUserLevelCommand_GivenOneAndAFailingRepository_WhenRunning_ThenRepositoryError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(userlevelcommand.TagUserLevelData, model2.UserLevelData{
		Code: "code",
	})
	ctx.Set(levelcommand.TagLevelId, uint(1000))
	ctx.Set(command.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	expectedUserLevel := model2.UserLevel{
		UserLevelData: model2.UserLevelData{
			Code: "code",
		},
		LevelId: uint(1000),
		UserId:  uint(1000),
	}
	s := new(repository.MockUserLevelRepository)
	s.On("StoreUserLevel", expectedUserLevel).Return(expectedUserLevel, expectedErr)
	cmd := userlevelcommand.CreateWriteUserLevelCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestWriteUserLevelCommand_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(userlevelcommand.TagUserLevelData, model2.UserLevelData{
		Code: "code",
	})
	ctx.Set(levelcommand.TagLevelId, uint(1000))
	ctx.Set(command.TagUserId, uint(1000))
	expectedVal := model2.UserLevel{
		UserLevelData: model2.UserLevelData{
			Code: "code",
		},
		LevelId: uint(1000),
		UserId:  uint(1000),
	}
	s := new(repository.MockUserLevelRepository)
	s.On("StoreUserLevel", expectedVal).Return(expectedVal, nil)
	cmd := userlevelcommand.CreateWriteUserLevelCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(userlevelcommand.TagUserLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
