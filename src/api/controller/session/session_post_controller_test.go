package session_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/executor/blocking"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/session"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	service2 "github.com/arpb2/C-3PO/src/api/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func createPostController() controller.Controller {
	return session.CreatePostController(
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

	reader := new(http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		obj.User = &model.User{
			Email: "test@email.com",
		}
		obj.Password = "test password"
		return true
	})).Return(nil).Once()

	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusBadRequest, http_wrapper.Json{
		"error": err.Error(),
	}).Once()
	middleware.On("IsAborted").Return(false).Once()
	middleware.On("IsAborted").Return(true).Once()

	validations := []session_validation.Validation{
		func(user *model.AuthenticatedUser) error {
			return nil
		},
		func(user *model.AuthenticatedUser) error {
			return err
		},
	}

	postController := session.CreatePostController(
		blocking.CreateExecutor(),
		nil,
		nil,
		validations,
	)

	postController.Body(&http_wrapper.Context{
		Reader:     reader,
		Writer:     nil,
		Middleware: middleware,
	})

	reader.AssertExpectations(t)
	middleware.AssertExpectations(t)
}

func TestFetchUserIdTaskImpl_FailsOnServiceFailure(t *testing.T) {
	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("IsAborted").Return(false).Times(3)
	middleware.On("AbortTransactionWithStatus", http.StatusInternalServerError, http_wrapper.Json{
		"error": "internal error",
	}).Once()

	reader := new (http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		obj.User = &model.User{
				Email: "test@email.com",
			}
		obj.Password = "testpassword"
		return true
	})).Return(nil).Once()

	service := new(service2.MockCredentialService)
	service.On("Retrieve", "test@email.com", "testpassword").Return(uint(0), errors.New("error")).Once()

	var validations []session_validation.Validation

	postController := session.CreatePostController(
		blocking.CreateExecutor(),
		nil,
		service,
		validations,
	)

	postController.Body(&http_wrapper.Context{
		Reader:     reader,
		Writer:     nil,
		Middleware: middleware,
	})

	middleware.AssertExpectations(t)
	service.AssertExpectations(t)
	reader.AssertExpectations(t)
}

func TestFetchUserIdTaskImpl_FailsOnTokenFailure(t *testing.T) {
	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusInternalServerError, http_wrapper.Json{
		"error": "internal error",
	}).Once()
	middleware.On("IsAborted").Return(false).Times(3)
	middleware.On("IsAborted").Return(true).Once()

	reader := new (http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		obj.User = &model.User{
			Email: "test@email.com",
		}
		obj.Password = "testpassword"
		return true
	})).Return(nil).Once()

	var validations []session_validation.Validation

	credentialService := new(service2.MockCredentialService)
	credentialService.On("Retrieve", "test@email.com", "testpassword").Return(uint(1000), nil)

	tokenHandler := new(auth.MockTokenHandler)
	tokenHandler.On("Create", mock.MatchedBy(func(tkn *auth.Token) bool {
		return tkn.UserId == uint(1000)
	})).Return("", http_wrapper.CreateInternalError())

	postController := session.CreatePostController(
		blocking.CreateExecutor(),
		tokenHandler,
		credentialService,
		validations,
	)

	postController.Body(&http_wrapper.Context{
		Reader:     reader,
		Writer:     nil,
		Middleware: middleware,
	})

	middleware.AssertExpectations(t)
	reader.AssertExpectations(t)
	credentialService.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func TestFetchUserIdTaskImpl_SuccessReturnsToken(t *testing.T) {
	writer := new(http_wrapper.MockWriter)
	writer.On("WriteJson", http.StatusOK, http_wrapper.Json{
		"token": "test token",
	}).Once()

	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("IsAborted").Return(false).Times(4)

	reader := new (http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		obj.User = &model.User{
			Email: "test@email.com",
		}
		obj.Password = "test password"
		return true
	})).Return(nil).Once()

	var validations []session_validation.Validation

	credentialService := new(service2.MockCredentialService)
	credentialService.On("Retrieve", "test@email.com", "test password").Return(uint(1000), nil)

	tokenHandler := new(auth.MockTokenHandler)
	tokenHandler.On("Create", mock.MatchedBy(func(tkn *auth.Token) bool {
		return tkn.UserId == uint(1000)
	})).Return("test token", nil)

	postController := session.CreatePostController(
		blocking.CreateExecutor(),
		tokenHandler,
		credentialService,
		validations,
	)

	postController.Body(&http_wrapper.Context{
		Reader:     reader,
		Writer:     writer,
		Middleware: middleware,
	})

	writer.AssertExpectations(t)
	reader.AssertExpectations(t)
	middleware.AssertExpectations(t)
	credentialService.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}