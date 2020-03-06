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

func TestGetLevelUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := level.CreateGetLevelUseCase(nil)

	name := useCase.Name()

	assert.Equal(t, "get_level_usecase", name)
}

func TestGetLevelUseCase_GivenOneAndAContextWithoutAuthenticatedLevel_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := level.CreateGetLevelUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestGetLevelUseCase_GivenOneAndAFailingRepository_WhenRunning_ThenRepositoryError(t *testing.T) {
	expectedVal := uint(1000)
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevelId, expectedVal)
	expectedErr := errors.New("some error")
	s := new(repository.MockLevelRepository)
	s.On("GetLevel", expectedVal).Return(level2.Level{}, expectedErr)
	useCase := level.CreateGetLevelUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetLevelUseCase_GivenOneAndANoLevelCreatedRepository_WhenRunning_Then404(t *testing.T) {
	expectedVal := uint(1000)
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevelId, expectedVal)
	expectedErr := http.CreateNotFoundError()
	s := new(repository.MockLevelRepository)
	s.On("GetLevel", expectedVal).Return(level2.Level{}, http.CreateNotFoundError())
	useCase := level.CreateGetLevelUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetLevelUseCase_GivenOne_WhenRunning_ThenContextHasLevelAndReturnsNoError(t *testing.T) {
	expectedVal := level2.Level{}
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevelId, uint(1000))
	s := new(repository.MockLevelRepository)
	s.On("GetLevel", uint(1000)).Return(expectedVal, nil)
	useCase := level.CreateGetLevelUseCase(s)

	err := useCase.Run(ctx)
	val, exists := ctx.Get(level.TagLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
