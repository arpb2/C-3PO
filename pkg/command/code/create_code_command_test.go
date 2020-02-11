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

func TestCreateCodeCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := code.CreateCreateCodeCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "create_code_command", name)
}

func TestCreateCodeCommand_GivenOneAndAContextWithoutCode_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	cmd := code.CreateCreateCodeCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestCreateCodeCommand_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(code.TagCodeRaw, "")
	cmd := code.CreateCreateCodeCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestCreateCodeCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(code.TagCodeRaw, "code")
	ctx.Set(user.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	s := new(service.MockCodeService)
	s.On("CreateCode", uint(1000), "code").Return(nil, expectedErr)
	cmd := code.CreateCreateCodeCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestCreateCodeCommand_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(code.TagCodeRaw, "code")
	ctx.Set(user.TagUserId, uint(1000))
	expectedVal := model.Code{}
	s := new(service.MockCodeService)
	s.On("CreateCode", uint(1000), "code").Return(&expectedVal, nil)
	cmd := code.CreateCreateCodeCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(code.TagCode)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
