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

func TestReplaceUserLevelCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := user_level.CreateReplaceUserLevelCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "replace_user_level_command", name)
}

func TestReplaceUserLevelCommand_GivenOneAndAContextWithoutRawCode_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	ctx.Set(user_level.TagLevelId, uint(1000))
	cmd := user_level.CreateReplaceUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestReplaceUserLevelCommand_GivenOneAndAContextWithoutLevelId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	ctx.Set(user_level.TagCodeRaw, "code")
	cmd := user_level.CreateReplaceUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestReplaceUserLevelCommand_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user_level.TagCodeRaw, "")
	ctx.Set(user_level.TagLevelId, uint(1000))
	cmd := user_level.CreateReplaceUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestReplaceUserLevelCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user_level.TagCodeRaw, "code")
	ctx.Set(user_level.TagLevelId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	expectedCode := model.UserLevel{
		Code:    "code",
		LevelId: uint(1000),
		UserId:  uint(1000),
	}
	s := new(service.MockUserLevelService)
	s.On("ReplaceUserLevel", &expectedCode).Return(expectedErr)
	cmd := user_level.CreateReplaceUserLevelCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestReplaceUserLevelCommand_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user_level.TagCodeRaw, "code")
	ctx.Set(user_level.TagLevelId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	expectedVal := model.UserLevel{
		Code:    "code",
		LevelId: uint(1000),
		UserId:  uint(1000),
	}
	s := new(service.MockUserLevelService)
	s.On("ReplaceUserLevel", &expectedVal).Return(nil)
	cmd := user_level.CreateReplaceUserLevelCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(user_level.TagUserLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
