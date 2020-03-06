package session_test

import (
	"errors"
	"testing"

	session3 "github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/arpb2/C-3PO/pkg/data/usecase/session"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	session2 "github.com/arpb2/C-3PO/pkg/domain/model/session"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/test/mock/auth"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestCreateTokenUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := session.CreateCreateSessionUseCase(nil)

	name := useCase.Name()

	assert.Equal(t, "create_session_usecase", name)
}

func TestCreateTokenUseCase_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := session.CreateCreateSessionUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestCreateTokenUseCase_GivenOneAndAFailingRepository_WhenRunning_ThenRepositoryError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	s := new(auth.MockTokenHandler)
	s.On("Create", &session3.Token{
		UserId: uint(1000),
	}).Return("", expectedErr)
	useCase := session.CreateCreateSessionUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestCreateTokenUseCase_GivenOne_WhenRunning_ThenContextHasUserIDAndReturnsNoError(t *testing.T) {
	expectedVal := session2.Session{
		UserId: uint(1000),
		Token:  "token",
	}
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, expectedVal.UserId)
	s := new(auth.MockTokenHandler)
	s.On("Create", &session3.Token{
		UserId: expectedVal.UserId,
	}).Return(expectedVal.Token, nil)
	useCase := session.CreateCreateSessionUseCase(s)

	err := useCase.Run(ctx)
	val, exists := ctx.Get(session.TagSession)

	assert.Nil(t, err)
	assert.Equal(t, expectedVal, val)
	assert.True(t, exists)
	s.AssertExpectations(t)
}
