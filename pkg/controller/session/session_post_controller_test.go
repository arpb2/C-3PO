package session_controller_test

import (
	"errors"
	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/controller/session"
	"github.com/arpb2/C-3PO/pkg/executor"
	user_validation "github.com/arpb2/C-3PO/pkg/validation/user"
	test_auth "github.com/arpb2/C-3PO/hack/auth"
	test_http_wrapper "github.com/arpb2/C-3PO/hack/http_wrapper"
	"github.com/arpb2/C-3PO/hack/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func createPostController() controller.Controller {
	return session_controller.CreatePostController(
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

	reader := new(test_http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		obj.User = &model.User{
			Email: "test@email.com",
		}
		obj.Password = "test password"
		return true
	})).Return(nil).Once()

	middleware := new(test_http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithError", http_wrapper.CreateBadRequestError(err.Error())).Once()

	validations := []user_validation.Validation{
		func(user *model.AuthenticatedUser) error {
			return nil
		},
		func(user *model.AuthenticatedUser) error {
			return err
		},
	}

	postController := session_controller.CreatePostController(
		executor.CreatePipeline(executor.CreateDebugHttpExecutor()),
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
	middleware := new(test_http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithError", http_wrapper.CreateInternalError()).Once()

	reader := new(test_http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		obj.User = &model.User{
			Email: "test@email.com",
		}
		obj.Password = "testpassword"
		return true
	})).Return(nil).Once()

	service := new(service.MockCredentialService)
	service.On("Retrieve", "test@email.com", "testpassword").Return(uint(0), http_wrapper.CreateInternalError()).Once()

	var validations []user_validation.Validation

	postController := session_controller.CreatePostController(
		executor.CreatePipeline(executor.CreateDebugHttpExecutor()),
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
	middleware := new(test_http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithError", http_wrapper.CreateInternalError()).Once()

	reader := new(test_http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		obj.User = &model.User{
			Email: "test@email.com",
		}
		obj.Password = "testpassword"
		return true
	})).Return(nil).Once()

	var validations []user_validation.Validation

	credentialService := new(service.MockCredentialService)
	credentialService.On("Retrieve", "test@email.com", "testpassword").Return(uint(1000), nil)

	tokenHandler := new(test_auth.MockTokenHandler)
	tokenHandler.On("Create", mock.MatchedBy(func(tkn *auth.Token) bool {
		return tkn.UserId == uint(1000)
	})).Return("", http_wrapper.CreateInternalError())

	postController := session_controller.CreatePostController(
		executor.CreatePipeline(executor.CreateDebugHttpExecutor()),
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
	writer := new(test_http_wrapper.MockWriter)
	writer.On("WriteJson", http.StatusOK, http_wrapper.Json{
		"token": "test token",
	}).Once()

	middleware := new(test_http_wrapper.MockMiddleware)

	reader := new(test_http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj *model.AuthenticatedUser) bool {
		obj.User = &model.User{
			Email: "test@email.com",
		}
		obj.Password = "test password"
		return true
	})).Return(nil).Once()

	var validations []user_validation.Validation

	credentialService := new(service.MockCredentialService)
	credentialService.On("Retrieve", "test@email.com", "test password").Return(uint(1000), nil)

	tokenHandler := new(test_auth.MockTokenHandler)
	tokenHandler.On("Create", mock.MatchedBy(func(tkn *auth.Token) bool {
		return tkn.UserId == uint(1000)
	})).Return("test token", nil)

	postController := session_controller.CreatePostController(
		executor.CreatePipeline(executor.CreateDebugHttpExecutor()),
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
