package session_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/command/session"
	"github.com/arpb2/C-3PO/pkg/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestRenderSessionCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := session.CreateRenderSessionCommand()

	name := cmd.Name()

	assert.Equal(t, "render_session_command", name)
}

func TestRenderSessionCommand_GivenOneAndAContextWithoutWriter_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(session.TagSession, model.Session{})
	cmd := session.CreateRenderSessionCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderSessionCommand_GivenOneAndAContextWithoutSession_WhenRunning_Then500(t *testing.T) {
	writer := new(http2.MockWriter)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpWriter, writer)
	cmd := session.CreateRenderSessionCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderSessionCommand_GivenOne_WhenRunning_ThenRendersSession(t *testing.T) {
	expectedVal := model.Session{}
	writer := new(http2.MockWriter)
	writer.On("WriteJson", 200, expectedVal)
	ctx := gopipeline.CreateContext()
	ctx.Set(session.TagSession, expectedVal)
	ctx.Set(pipeline.TagHttpWriter, writer)
	cmd := session.CreateRenderSessionCommand()

	err := cmd.Run(ctx)

	assert.Nil(t, err)
	writer.AssertExpectations(t)
}
