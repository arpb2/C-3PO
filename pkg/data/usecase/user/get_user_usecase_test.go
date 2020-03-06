package user_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestGetUserUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := user.CreateGetUserUseCase(nil)

	name := useCase.Name()

	assert.Equal(t, "get_user_usecase", name)
}

func TestGetUserUseCase_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := user.CreateGetUserUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestGetUserUseCase_GivenOneAndAFailingRepository_WhenRunning_ThenRepositoryError(t *testing.T) {
	expectedVal := uint(1000)
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, expectedVal)
	expectedErr := errors.New("some error")
	s := new(repository.MockUserRepository)
	s.On("GetUser", expectedVal).Return(nil, expectedErr)
	useCase := user.CreateGetUserUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetUserUseCase_GivenOneAndANoUserCreatedRepository_WhenRunning_Then404(t *testing.T) {
	expectedVal := uint(1000)
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, expectedVal)
	expectedErr := http.CreateNotFoundError()
	s := new(repository.MockUserRepository)
	s.On("GetUser", expectedVal).Return(user2.User{}, http.CreateNotFoundError())
	useCase := user.CreateGetUserUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetUserUseCase_GivenOne_WhenRunning_ThenContextHasUserAndReturnsNoError(t *testing.T) {
	expectedVal := user2.User{}
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	s := new(repository.MockUserRepository)
	s.On("GetUser", uint(1000)).Return(expectedVal, nil)
	useCase := user.CreateGetUserUseCase(s)

	err := useCase.Run(ctx)
	val, exists := ctx.Get(user.TagUser)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
