package session_test

import (
	"errors"
	"net/http"
	"testing"

	session3 "github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/domain/model/session"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
	session2 "github.com/arpb2/C-3PO/pkg/presentation/session"

	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"

	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
	testauth "github.com/arpb2/C-3PO/test/mock/auth"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	mockrepository "github.com/arpb2/C-3PO/test/mock/repository"
	"github.com/stretchr/testify/mock"
)

func TestPostController_FetchUserIdTask_FailsOnValidationFail(t *testing.T) {
	err := errors.New("second throws error")

	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *user.AuthenticatedUser) bool {
		obj.User = user.User{
			Email: "test@email.com",
		}
		obj.Password = "test password"
		return true
	})).Return(nil).Once()

	middleware := new(testhttpwrapper.MockMiddleware)
	middleware.On("AbortTransactionWithError", httpwrapper.CreateBadRequestError(err.Error())).Once()

	validations := []validation.Validation{
		func(user *user.AuthenticatedUser) error {
			return nil
		},
		func(user *user.AuthenticatedUser) error {
			return err
		},
	}

	handler := session2.CreatePostHandler(
		pipeline2.CreateDebugHttpPipeline(),
		nil,
		nil,
		validations,
	)

	handler(&httpwrapper.Context{
		Reader:     reader,
		Writer:     nil,
		Middleware: middleware,
	})

	reader.AssertExpectations(t)
	middleware.AssertExpectations(t)
}

func TestFetchUserIdTaskImpl_FailsOnRepositoryFailure(t *testing.T) {
	middleware := new(testhttpwrapper.MockMiddleware)
	middleware.On("AbortTransactionWithError", httpwrapper.CreateInternalError()).Once()

	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *user.AuthenticatedUser) bool {
		obj.User = user.User{
			Email: "test@email.com",
		}
		obj.Password = "testpassword"
		return true
	})).Return(nil).Once()

	repository := new(mockrepository.MockCredentialRepository)
	repository.On("GetUserId", "test@email.com", "testpassword").Return(uint(0), httpwrapper.CreateInternalError()).Once()

	var validations []validation.Validation

	handler := session2.CreatePostHandler(
		pipeline2.CreateDebugHttpPipeline(),
		nil,
		repository,
		validations,
	)

	handler(&httpwrapper.Context{
		Reader:     reader,
		Writer:     nil,
		Middleware: middleware,
	})

	middleware.AssertExpectations(t)
	repository.AssertExpectations(t)
	reader.AssertExpectations(t)
}

func TestFetchUserIdTaskImpl_FailsOnTokenFailure(t *testing.T) {
	middleware := new(testhttpwrapper.MockMiddleware)
	middleware.On("AbortTransactionWithError", httpwrapper.CreateInternalError()).Once()

	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *user.AuthenticatedUser) bool {
		obj.User = user.User{
			Email: "test@email.com",
		}
		obj.Password = "testpassword"
		return true
	})).Return(nil).Once()

	var validations []validation.Validation

	credentialRepository := new(mockrepository.MockCredentialRepository)
	credentialRepository.On("GetUserId", "test@email.com", "testpassword").Return(uint(1000), nil)

	tokenHandler := new(testauth.MockTokenHandler)
	tokenHandler.On("Create", mock.MatchedBy(func(tkn *session3.Token) bool {
		return tkn.UserId == uint(1000)
	})).Return("", httpwrapper.CreateInternalError())

	handler := session2.CreatePostHandler(
		pipeline2.CreateDebugHttpPipeline(),
		tokenHandler,
		credentialRepository,
		validations,
	)

	handler(&httpwrapper.Context{
		Reader:     reader,
		Writer:     nil,
		Middleware: middleware,
	})

	middleware.AssertExpectations(t)
	reader.AssertExpectations(t)
	credentialRepository.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func TestFetchUserIdTaskImpl_SuccessReturnsToken(t *testing.T) {
	writer := new(testhttpwrapper.MockWriter)
	writer.On("WriteJson", http.StatusOK, session.Session{
		UserId: uint(1000),
		Token:  "test token",
	}).Once()

	middleware := new(testhttpwrapper.MockMiddleware)

	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *user.AuthenticatedUser) bool {
		obj.User = user.User{
			Email: "test@email.com",
		}
		obj.Password = "test password"
		return true
	})).Return(nil).Once()

	var validations []validation.Validation

	credentialRepository := new(mockrepository.MockCredentialRepository)
	credentialRepository.On("GetUserId", "test@email.com", "test password").Return(uint(1000), nil)

	tokenHandler := new(testauth.MockTokenHandler)
	tokenHandler.On("Create", mock.MatchedBy(func(tkn *session3.Token) bool {
		return tkn.UserId == uint(1000)
	})).Return("test token", nil)

	handler := session2.CreatePostHandler(
		pipeline2.CreateDebugHttpPipeline(),
		tokenHandler,
		credentialRepository,
		validations,
	)

	handler(&httpwrapper.Context{
		Reader:     reader,
		Writer:     writer,
		Middleware: middleware,
	})

	writer.AssertExpectations(t)
	reader.AssertExpectations(t)
	middleware.AssertExpectations(t)
	credentialRepository.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}
