package level_test

import (
	"fmt"
	"testing"

	level2 "github.com/arpb2/C-3PO/pkg/data/usecase/level"

	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestFetchLevelIdUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := level2.CreateFetchLevelIdUseCase("level_id")

	name := useCase.Name()

	assert.Equal(t, "fetch_level_id_usecase", name)
}

func TestFetchLevelIdUseCase_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := level2.CreateFetchLevelIdUseCase("level_id")

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchLevelIdUseCase_GivenOneAndAReaderWithoutLevelIdParameter_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", "level_id").Return("", false)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	useCase := level2.CreateFetchLevelIdUseCase("level_id")

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(fmt.Sprintf("'%s' empty", "level_id")), err)
	reader.AssertExpectations(t)
}

func TestFetchLevelIdUseCase_GivenOneAndAReaderWithMalformedLevelId_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", "level_id").Return("-1", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	useCase := level2.CreateFetchLevelIdUseCase("level_id")

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", "level_id")), err)
	reader.AssertExpectations(t)
}

func TestFetchLevelIdUseCase_GivenOne_WhenRunning_ThenRawLevelIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", "level_id").Return("1000", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	useCase := level2.CreateFetchLevelIdUseCase("level_id")

	err := useCase.Run(ctx)
	val, exists := ctx.Get(level2.TagLevelId)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, uint(1000), val)
	reader.AssertExpectations(t)
}
