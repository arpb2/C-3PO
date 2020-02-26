package command_test

import (
	"testing"

	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestRenderEmptyCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateRenderEmptyCommand()

	name := cmd.Name()

	assert.Equal(t, "render_empty_command", name)
}

func TestRenderEmptyCommand_GivenOneAndAContextWithoutWriter_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateRenderEmptyCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderEmptyCommand_GivenOne_WhenRunning_ThenRendersUser(t *testing.T) {
	writer := new(http2.MockWriter)
	writer.On("WriteStatus", 200)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpWriter, writer)
	cmd := command.CreateRenderEmptyCommand()

	err := cmd.Run(ctx)

	assert.Nil(t, err)
	writer.AssertExpectations(t)
}
