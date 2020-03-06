package user_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestRenderUserLevelUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := user.CreateRenderUserLevelUseCase()

	name := useCase.Name()

	assert.Equal(t, "render_user_level_usecase", name)
}

func TestRenderUserLevelUseCase_GivenOneAndAContextWithoutWriter_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserLevel, user2.Level{})
	useCase := user.CreateRenderUserLevelUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderUserLevelUseCase_GivenOneAndAContextWithoutCode_WhenRunning_Then500(t *testing.T) {
	writer := new(http2.MockWriter)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpWriter, writer)
	useCase := user.CreateRenderUserLevelUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderUserLevelUseCase_GivenOne_WhenRunning_ThenRendersCode(t *testing.T) {
	expectedVal := user2.Level{}
	writer := new(http2.MockWriter)
	writer.On("WriteJson", 200, expectedVal)
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserLevel, expectedVal)
	ctx.Set(pipeline2.TagHttpWriter, writer)
	useCase := user.CreateRenderUserLevelUseCase()

	err := useCase.Run(ctx)

	assert.Nil(t, err)
	writer.AssertExpectations(t)
}
