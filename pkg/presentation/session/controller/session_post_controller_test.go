package controller_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/session/repository"

	model3 "github.com/arpb2/C-3PO/pkg/domain/session/model"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
	pipeline2 "github.com/arpb2/C-3PO/test/mock/pipeline"

	sessioncontroller "github.com/arpb2/C-3PO/pkg/presentation/session/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	testauth "github.com/arpb2/C-3PO/test/mock/auth"
	testhttpwrapper "github.com/arpb2/C-3PO/test/mock/http"
	mockrepository "github.com/arpb2/C-3PO/test/mock/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createPostController() controller.Controller {
	return sessioncontroller.CreatePostController(
		nil,
		nil,
		nil,
		nil,
	)
}

func TestPostController_IsPost(t *testing.T) {
	assert.Equal(t, "POST", createPostController().Method)
}

func TestPostControllerPath_IsSession(t *testing.T) {
	assert.Equal(t, "/session", createPostController().Path)
}

func TestPostController_FetchUserIdTask_FailsOnValidationFail(t *testing.T) {
	err := errors.New("second throws error")

	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model2.AuthenticatedUser) bool {
		obj.User = model2.User{
			Email: "test@email.com",
		}
		obj.Password = "test password"
		return true
	})).Return(nil).Once()

	middleware := new(testhttpwrapper.MockMiddleware)
	middleware.On("AbortTransactionWithError", httpwrapper.CreateBadRequestError(err.Error())).Once()

	validations := []validation.Validation{
		func(user *model2.AuthenticatedUser) error {
			return nil
		},
		func(user *model2.AuthenticatedUser) error {
			return err
		},
	}

	postController := sessioncontroller.CreatePostController(
		pipeline2.CreateDebugHttpPipeline(),
		nil,
		nil,
		validations,
	)

	postController.Body(&httpwrapper.Context{
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
	reader.On("ReadBody", mock.MatchedBy(func(obj *model2.AuthenticatedUser) bool {
		obj.User = model2.User{
			Email: "test@email.com",
		}
		obj.Password = "testpassword"
		return true
	})).Return(nil).Once()

	repository := new(mockrepository.MockCredentialRepository)
	repository.On("GetUserId", "test@email.com", "testpassword").Return(uint(0), httpwrapper.CreateInternalError()).Once()

	var validations []validation.Validation

	postController := sessioncontroller.CreatePostController(
		pipeline2.CreateDebugHttpPipeline(),
		nil,
		repository,
		validations,
	)

	postController.Body(&httpwrapper.Context{
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
	reader.On("ReadBody", mock.MatchedBy(func(obj *model2.AuthenticatedUser) bool {
		obj.User = model2.User{
			Email: "test@email.com",
		}
		obj.Password = "testpassword"
		return true
	})).Return(nil).Once()

	var validations []validation.Validation

	credentialRepository := new(mockrepository.MockCredentialRepository)
	credentialRepository.On("GetUserId", "test@email.com", "testpassword").Return(uint(1000), nil)

	tokenHandler := new(testauth.MockTokenHandler)
	tokenHandler.On("Create", mock.MatchedBy(func(tkn *repository.Token) bool {
		return tkn.UserId == uint(1000)
	})).Return("", httpwrapper.CreateInternalError())

	postController := sessioncontroller.CreatePostController(
		pipeline2.CreateDebugHttpPipeline(),
		tokenHandler,
		credentialRepository,
		validations,
	)

	postController.Body(&httpwrapper.Context{
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
	writer.On("WriteJson", http.StatusOK, model3.Session{
		UserId: uint(1000),
		Token:  "test token",
	}).Once()

	middleware := new(testhttpwrapper.MockMiddleware)

	reader := new(testhttpwrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model2.AuthenticatedUser) bool {
		obj.User = model2.User{
			Email: "test@email.com",
		}
		obj.Password = "test password"
		return true
	})).Return(nil).Once()

	var validations []validation.Validation

	credentialRepository := new(mockrepository.MockCredentialRepository)
	credentialRepository.On("GetUserId", "test@email.com", "test password").Return(uint(1000), nil)

	tokenHandler := new(testauth.MockTokenHandler)
	tokenHandler.On("Create", mock.MatchedBy(func(tkn *repository.Token) bool {
		return tkn.UserId == uint(1000)
	})).Return("test token", nil)

	postController := sessioncontroller.CreatePostController(
		pipeline2.CreateDebugHttpPipeline(),
		tokenHandler,
		credentialRepository,
		validations,
	)

	postController.Body(&httpwrapper.Context{
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
