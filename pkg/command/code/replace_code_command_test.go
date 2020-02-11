package code_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/command/code"
	"github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestReplaceCodeCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := code.CreateReplaceCodeCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "replace_code_command", name)
}

func TestReplaceCodeCommand_GivenOneAndAContextWithoutRawCode_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	ctx.Set(code.TagCodeId, uint(1000))
	cmd := code.CreateReplaceCodeCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestReplaceCodeCommand_GivenOneAndAContextWithoutCodeId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	ctx.Set(code.TagCodeRaw, "code")
	cmd := code.CreateReplaceCodeCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestReplaceCodeCommand_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(code.TagCodeRaw, "")
	ctx.Set(code.TagCodeId, uint(1000))
	cmd := code.CreateReplaceCodeCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestReplaceCodeCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(code.TagCodeRaw, "code")
	ctx.Set(code.TagCodeId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	expectedCode := model.Code{
		Code:   "code",
		Id:     uint(1000),
		UserId: uint(1000),
	}
	s := new(service.MockCodeService)
	s.On("ReplaceCode", &expectedCode).Return(expectedErr)
	cmd := code.CreateReplaceCodeCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestReplaceCodeCommand_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(code.TagCodeRaw, "code")
	ctx.Set(code.TagCodeId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	expectedVal := model.Code{
		Code:   "code",
		Id:     uint(1000),
		UserId: uint(1000),
	}
	s := new(service.MockCodeService)
	s.On("ReplaceCode", &expectedVal).Return(nil)
	cmd := code.CreateReplaceCodeCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(code.TagCode)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
