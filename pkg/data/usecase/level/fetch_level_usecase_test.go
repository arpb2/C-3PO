package level_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/level"
	level2 "github.com/arpb2/C-3PO/pkg/domain/model/level"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchLevelUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := level.CreateFetchLevelUseCase()

	name := useCase.Name()

	assert.Equal(t, "fetch_level_usecase", name)
}

func TestFetchLevelUseCase_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := level.CreateFetchLevelUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchLevelUseCase_GivenOneAndAReaderWithoutBody_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("ReadBody", mock.Anything).Return(errors.New("some error"))
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	useCase := level.CreateFetchLevelUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("malformed body"), err)
	reader.AssertExpectations(t)
}

func TestFetchLevelUseCase_GivenOne_WhenRunning_ThenLevelIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("ReadBody", mock.Anything).Return(nil)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	useCase := level.CreateFetchLevelUseCase()

	err := useCase.Run(ctx)
	val, exists := ctx.Get(level.TagLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, level2.Level{}, val)
	reader.AssertExpectations(t)
}
