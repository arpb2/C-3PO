package level_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/level"
	level2 "github.com/arpb2/C-3PO/pkg/domain/model/level"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestRenderLevelUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := level.CreateRenderLevelUseCase()

	name := useCase.Name()

	assert.Equal(t, "render_level_usecase", name)
}

func TestRenderLevelUseCase_GivenOneAndAContextWithoutWriter_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevel, level2.Level{})
	useCase := level.CreateRenderLevelUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderLevelUseCase_GivenOneAndAContextWithoutLevel_WhenRunning_Then500(t *testing.T) {
	writer := new(http2.MockWriter)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpWriter, writer)
	useCase := level.CreateRenderLevelUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderLevelUseCase_GivenOne_WhenRunning_ThenRendersLevel(t *testing.T) {
	expectedVal := level2.Level{}
	writer := new(http2.MockWriter)
	writer.On("WriteJson", 200, expectedVal)
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevel, level2.Level{})
	ctx.Set(pipeline2.TagHttpWriter, writer)
	useCase := level.CreateRenderLevelUseCase()

	err := useCase.Run(ctx)

	assert.Nil(t, err)
	writer.AssertExpectations(t)
}
