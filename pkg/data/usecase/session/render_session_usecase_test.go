package session_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/session"
	session2 "github.com/arpb2/C-3PO/pkg/domain/model/session"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestRenderSessionUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := session.CreateRenderSessionUseCase()

	name := useCase.Name()

	assert.Equal(t, "render_session_usecase", name)
}

func TestRenderSessionUseCase_GivenOneAndAContextWithoutWriter_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(session.TagSession, session2.Session{})
	useCase := session.CreateRenderSessionUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderSessionUseCase_GivenOneAndAContextWithoutSession_WhenRunning_Then500(t *testing.T) {
	writer := new(http2.MockWriter)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpWriter, writer)
	useCase := session.CreateRenderSessionUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderSessionUseCase_GivenOne_WhenRunning_ThenRendersSession(t *testing.T) {
	expectedVal := session2.Session{}
	writer := new(http2.MockWriter)
	writer.On("WriteJson", 200, expectedVal)
	ctx := gopipeline.CreateContext()
	ctx.Set(session.TagSession, expectedVal)
	ctx.Set(pipeline2.TagHttpWriter, writer)
	useCase := session.CreateRenderSessionUseCase()

	err := useCase.Run(ctx)

	assert.Nil(t, err)
	writer.AssertExpectations(t)
}
