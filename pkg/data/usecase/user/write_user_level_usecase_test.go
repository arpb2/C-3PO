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

func TestWriteUserLevelUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := user.CreateWriteUserLevelUseCase(nil)

	name := useCase.Name()

	assert.Equal(t, "write_user_level_usecase", name)
}

func TestWriteUserLevelUseCase_GivenOneAndAContextWithoutRawCode_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	ctx.Set(level.TagLevelId, uint(1000))
	useCase := user.CreateWriteUserLevelUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteUserLevelUseCase_GivenOneAndAContextWithoutLevelId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	ctx.Set(user.TagUserLevelData, user2.LevelData{
		Code: "code",
	})
	useCase := user.CreateWriteUserLevelUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteUserLevelUseCase_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserLevelData, user2.LevelData{})
	ctx.Set(level.TagLevelId, uint(1000))
	useCase := user.CreateWriteUserLevelUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestWriteUserLevelUseCase_GivenOneAndAFailingRepository_WhenRunning_ThenRepositoryError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserLevelData, user2.LevelData{
		Code: "code",
	})
	ctx.Set(level.TagLevelId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	expectedUserLevel := user2.Level{
		LevelData: user2.LevelData{
			Code: "code",
		},
		LevelId: uint(1000),
		UserId:  uint(1000),
	}
	s := new(repository.MockUserLevelRepository)
	s.On("StoreUserLevel", expectedUserLevel).Return(expectedUserLevel, expectedErr)
	useCase := user.CreateWriteUserLevelUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestWriteUserLevelUseCase_GivenOne_WhenRunning_ThenContextHasCodeAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserLevelData, user2.LevelData{
		Code: "code",
	})
	ctx.Set(level.TagLevelId, uint(1000))
	ctx.Set(user.TagUserId, uint(1000))
	expectedVal := user2.Level{
		LevelData: user2.LevelData{
			Code: "code",
		},
		LevelId: uint(1000),
		UserId:  uint(1000),
	}
	s := new(repository.MockUserLevelRepository)
	s.On("StoreUserLevel", expectedVal).Return(expectedVal, nil)
	useCase := user.CreateWriteUserLevelUseCase(s)

	err := useCase.Run(ctx)
	val, exists := ctx.Get(user.TagUserLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
