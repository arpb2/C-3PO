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

func TestCreateUserLevelCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := user_level.CreateCreateUserLevelCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "create_user_level_command", name)
}

func TestCreateUserLevelCommand_GivenOneAndAContextWithoutCode_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	cmd := user_level.CreateCreateUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestCreateUserLevelCommand_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user_level.TagCodeRaw, "")
	cmd := user_level.CreateCreateUserLevelCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestCreateUserLevelCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user_level.TagCodeRaw, "code")
	ctx.Set(user.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	s := new(service.MockUserLevelService)
	s.On("CreateUserLevel", uint(1000), "code").Return(nil, expectedErr)
	cmd := user_level.CreateCreateUserLevelCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestCreateUserLevelCommand_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user_level.TagCodeRaw, "code")
	ctx.Set(user.TagUserId, uint(1000))
	expectedVal := model.UserLevel{}
	s := new(service.MockUserLevelService)
	s.On("CreateUserLevel", uint(1000), "code").Return(&expectedVal, nil)
	cmd := user_level.CreateCreateUserLevelCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(user_level.TagUserLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
