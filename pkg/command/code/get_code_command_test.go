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

func TestGetCodeCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := code.CreateGetCodeCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "get_code_command", name)
}

func TestGetCodeCommand_GivenOneAndAContextWithoutCodeID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	cmd := code.CreateGetCodeCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestGetCodeCommand_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(code.TagCodeId, uint(1000))
	cmd := code.CreateGetCodeCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestGetCodeCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(code.TagCodeId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	s := new(service.MockCodeService)
	s.On("GetCode", uint(1000), uint(1000)).Return(nil, expectedErr)
	cmd := code.CreateGetCodeCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetCodeCommand_GivenOneAndAServiceWithNoCode_WhenRunning_Then404(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(code.TagCodeId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	s := new(service.MockCodeService)
	s.On("GetCode", uint(1000), uint(1000)).Return(nil, nil)
	cmd := code.CreateGetCodeCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateNotFoundError(), err)
	s.AssertExpectations(t)
}

func TestGetCodeCommand_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(code.TagCodeId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	expectedVal := model.Code{}
	s := new(service.MockCodeService)
	s.On("GetCode", uint(1000), uint(1000)).Return(&expectedVal, nil)
	cmd := code.CreateGetCodeCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(code.TagCode)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
