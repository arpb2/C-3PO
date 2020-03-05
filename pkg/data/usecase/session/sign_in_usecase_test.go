package session_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/session"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticateUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := session.CreateAuthenticateUseCase(nil)

	name := useCase.Name()

	assert.Equal(t, "authenticate_usecase", name)
}

func TestAuthenticateUseCase_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := session.CreateAuthenticateUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestAuthenticateUseCase_GivenOneAndAFailingRepository_WhenRunning_ThenRepositoryError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, user2.AuthenticatedUser{
		User: user2.User{},
	})
	expectedErr := errors.New("some error")
	s := new(repository.MockCredentialRepository)
	s.On("GetUserId", "", "").Return(uint(0), expectedErr)
	useCase := session.CreateAuthenticateUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestAuthenticateUseCase_GivenOne_WhenRunning_ThenContextHasUserIDAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, user2.AuthenticatedUser{
		User: user2.User{},
	})
	expectedVal := uint(1000)
	s := new(repository.MockCredentialRepository)
	s.On("GetUserId", "", "").Return(expectedVal, nil)
	useCase := session.CreateAuthenticateUseCase(s)

	err := useCase.Run(ctx)
	val, exists := ctx.Get(user.TagUserId)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
