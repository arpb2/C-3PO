package user_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/level"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestGetUserLevelUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := user.CreateGetUserLevelUseCase(nil)

	name := useCase.Name()

	assert.Equal(t, "get_user_level_usecase", name)
}

func TestGetUserLevelUseCase_GivenOneAndAContextWithoutLevelID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	useCase := user.CreateGetUserLevelUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestGetUserLevelUseCase_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevelId, uint(1000))
	useCase := user.CreateGetUserLevelUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestGetUserLevelUseCase_GivenOneAndAFailingRepository_WhenRunning_ThenRepositoryError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevelId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	s := new(repository.MockUserLevelRepository)
	s.On("GetUserLevel", uint(1000), uint(1000)).Return(user2.Level{}, expectedErr)
	useCase := user.CreateGetUserLevelUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetUserLevelUseCase_GivenOneAndARepositoryWithNoCode_WhenRunning_Then404(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevelId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	s := new(repository.MockUserLevelRepository)
	s.On("GetUserLevel", uint(1000), uint(1000)).Return(user2.Level{}, http.CreateNotFoundError())
	useCase := user.CreateGetUserLevelUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateNotFoundError(), err)
	s.AssertExpectations(t)
}

func TestGetUserLevelUseCase_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(level.TagLevelId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	expectedVal := user2.Level{}
	s := new(repository.MockUserLevelRepository)
	s.On("GetUserLevel", uint(1000), uint(1000)).Return(expectedVal, nil)
	useCase := user.CreateGetUserLevelUseCase(s)

	err := useCase.Run(ctx)
	val, exists := ctx.Get(user.TagUserLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
