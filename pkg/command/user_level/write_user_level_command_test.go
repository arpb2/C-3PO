package user_level_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/arpb2/C-3PO/pkg/command/user_level"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestWriteUserLevelCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := user_level.CreateWriteUserLevelCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "write_user_level_command", name)
}

func TestWriteUserLevelCommand_GivenOneAndAContextWithoutRawCode_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	ctx.Set(user_level.TagLevelId, uint(1000))
	cmd := user_level.CreateWriteUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteUserLevelCommand_GivenOneAndAContextWithoutLevelId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	ctx.Set(user_level.TagUserLevelData, model.UserLevelData{
		Code: "code",
	})
	cmd := user_level.CreateWriteUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteUserLevelCommand_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user_level.TagUserLevelData, model.UserLevelData{})
	ctx.Set(user_level.TagLevelId, uint(1000))
	cmd := user_level.CreateWriteUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteUserLevelCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user_level.TagUserLevelData, model.UserLevelData{
		Code: "code",
	})
	ctx.Set(user_level.TagLevelId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	expectedUserLevel := model.UserLevel{
		UserLevelData: &model.UserLevelData{
			Code: "code",
		},
		LevelId: uint(1000),
		UserId:  uint(1000),
	}
	s := new(service.MockUserLevelService)
	s.On("WriteUserLevel", expectedUserLevel).Return(expectedUserLevel, expectedErr)
	cmd := user_level.CreateWriteUserLevelCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestWriteUserLevelCommand_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user_level.TagUserLevelData, model.UserLevelData{
		Code: "code",
	})
	ctx.Set(user_level.TagLevelId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	expectedVal := model.UserLevel{
		UserLevelData: &model.UserLevelData{
			Code: "code",
		},
		LevelId: uint(1000),
		UserId:  uint(1000),
	}
	s := new(service.MockUserLevelService)
	s.On("WriteUserLevel", expectedVal).Return(expectedVal, nil)
	cmd := user_level.CreateWriteUserLevelCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(user_level.TagUserLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
