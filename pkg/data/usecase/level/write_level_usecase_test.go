package level_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/level"
	level2 "github.com/arpb2/C-3PO/pkg/domain/model/level"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestWriteUserLevelUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := level.CreateWriteLevelUseCase(nil)

	name := useCase.Name()

	assert.Equal(t, "write_level_usecase", name)
}

func TestWriteLevelUseCase_GivenOneAndAContextWithoutLevel_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevelId, uint(1000))
	useCase := level.CreateWriteLevelUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteLevelUseCase_GivenOneAndAContextWithoutLevelId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevel, level2.Level{})
	useCase := level.CreateWriteLevelUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteLevelUseCase_GivenOneAndAFailingRepository_WhenRunning_ThenRepositoryError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevel, level2.Level{})
	ctx.Set(level.TagLevelId, uint(1000))
	expectedErr := errors.New("some error")
	expectedLevel := level2.Level{}
	s := new(repository.MockLevelRepository)
	s.On("StoreLevel", level2.Level{
		Id: 1000,
	}).Return(expectedLevel, expectedErr)
	useCase := level.CreateWriteLevelUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestWriteLevelUseCase_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevel, level2.Level{})
	ctx.Set(level.TagLevelId, uint(1000))
	expectedVal := level2.Level{}
	s := new(repository.MockLevelRepository)
	s.On("StoreLevel", level2.Level{
		Id: 1000,
	}).Return(expectedVal, nil)
	useCase := level.CreateWriteLevelUseCase(s)

	err := useCase.Run(ctx)
	val, exists := ctx.Get(level.TagLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
