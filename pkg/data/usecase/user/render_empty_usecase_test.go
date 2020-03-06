package user_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user"

	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestRenderEmptyUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := user.CreateRenderEmptyUseCase()

	name := useCase.Name()

	assert.Equal(t, "render_empty_usecase", name)
}

func TestRenderEmptyUseCase_GivenOneAndAContextWithoutWriter_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := user.CreateRenderEmptyUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderEmptyUseCase_GivenOne_WhenRunning_ThenRendersUser(t *testing.T) {
	writer := new(http2.MockWriter)
	writer.On("WriteStatus", 200)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpWriter, writer)
	useCase := user.CreateRenderEmptyUseCase()

	err := useCase.Run(ctx)

	assert.Nil(t, err)
	writer.AssertExpectations(t)
}
