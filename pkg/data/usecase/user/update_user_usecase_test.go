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

func TestUpdateUserUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := user.CreateUpdateUserUseCase(nil)

	name := useCase.Name()

	assert.Equal(t, "update_user_usecase", name)
}

func TestUpdateUserUseCase_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	useCase := user.CreateUpdateUserUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestUpdateUserUseCase_GivenOneAndAContextWithoutUserId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, user2.AuthenticatedUser{})
	useCase := user.CreateUpdateUserUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestUpdateUserUseCase_GivenOneAndAFailingRepository_WhenRunning_ThenRepositoryError(t *testing.T) {
	expectedVal := user2.AuthenticatedUser{
		User: user2.User{
			Id: uint(1000),
		},
	}
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	ctx.Set(user.TagAuthenticatedUser, expectedVal)
	expectedErr := errors.New("some error")
	s := new(repository.MockUserRepository)
	s.On("UpdateUser", expectedVal).Return(expectedVal.User, expectedErr)
	useCase := user.CreateUpdateUserUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestUpdateUserUseCase_GivenOne_WhenRunning_ThenContextHasUserAndReturnsNoError(t *testing.T) {
	expectedVal := user2.AuthenticatedUser{
		User: user2.User{
			Id: uint(1000),
		},
	}
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	ctx.Set(user.TagAuthenticatedUser, expectedVal)
	s := new(repository.MockUserRepository)
	s.On("UpdateUser", expectedVal).Return(expectedVal.User, nil)
	useCase := user.CreateUpdateUserUseCase(s)

	err := useCase.Run(ctx)
	val, exists := ctx.Get(user.TagUser)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal.User, val)
	s.AssertExpectations(t)
}
