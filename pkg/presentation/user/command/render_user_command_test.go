package command_test

import (
	"testing"

	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestRenderUserCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateRenderUserCommand()

	name := cmd.Name()

	assert.Equal(t, "render_user_command", name)
}

func TestRenderUserCommand_GivenOneAndAContextWithoutWriter_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUser, model2.User{})
	cmd := command.CreateRenderUserCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderUserCommand_GivenOneAndAContextWithoutUser_WhenRunning_Then500(t *testing.T) {
	writer := new(http2.MockWriter)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpWriter, writer)
	cmd := command.CreateRenderUserCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderUserCommand_GivenOne_WhenRunning_ThenRendersUser(t *testing.T) {
	expectedVal := model2.User{}
	writer := new(http2.MockWriter)
	writer.On("WriteJson", 200, expectedVal)
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUser, model2.User{})
	ctx.Set(pipeline2.TagHttpWriter, writer)
	cmd := command.CreateRenderUserCommand()

	err := cmd.Run(ctx)

	assert.Nil(t, err)
	writer.AssertExpectations(t)
}
