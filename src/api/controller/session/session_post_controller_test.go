package session_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/session"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/arpb2/C-3PO/src/api/task/authenticated_user_task"
	"github.com/arpb2/C-3PO/src/api/task/token_task"
	"github.com/arpb2/C-3PO/src/api/validation/authenticated_user_validation"
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

	reader := new(http_wrapper.TestReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil).Once()

	middleware := new(http_wrapper.TestMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusBadRequest, http_wrapper.Json{
		"error": err.Error(),
	}).Once()

	validations := []authenticated_user_validation.Validation{
		func(user *model.AuthenticatedUser) error {
			return nil
		},
		func(user *model.AuthenticatedUser) error {
			return err
		},
	}
	fetchUserTask := authenticated_user_task.FetchUserTaskImpl

	postController := session.CreatePostController(
		nil,
		nil,
		validations,
		fetchUserTask,
		nil,
		nil,
	)

	postController.Body(&http_wrapper.Context{
		Reader:     reader,
		Writer:     nil,
		Middleware: middleware,
	})

	reader.AssertExpectations(t)
	middleware.AssertExpectations(t)
}

type credentialService struct{
	mock.Mock
}

func (c credentialService) Store(user *model.AuthenticatedUser) error {
	args := c.Called(user)
	return args.Error(0)
}

func (c credentialService) Retrieve(email, password string) (uint, error) {
	args := c.Called(email, password)
	return args.Get(0).(uint), args.Error(1)
}

func TestFetchUserIdTaskImpl_FailsOnServiceFailure(t *testing.T) {
	middleware := new(http_wrapper.TestMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusInternalServerError, http_wrapper.Json{
		"error": "internal error",
	}).Once()

	service := new(credentialService)
	service.On("Retrieve", "test@email.com", "testpassword").Return(uint(0), errors.New("error")).Once()

	var validations []authenticated_user_validation.Validation
	fetchUserTask := func(ctx *http_wrapper.Context) (user *model.AuthenticatedUser, err error) {
		return &model.AuthenticatedUser{
			User: &model.User{
				Email: "test@email.com",
			},
			Password: "testpassword",
		}, nil
	}
	fetchUserIdTask := authenticated_user_task.FetchUserIdTaskImpl

	postController := session.CreatePostController(
		nil,
		service,
		validations,
		fetchUserTask,
		fetchUserIdTask,
		nil,
	)

	postController.Body(&http_wrapper.Context{
		Reader:     nil,
		Writer:     nil,
		Middleware: middleware,
	})

	middleware.AssertExpectations(t)
	service.AssertExpectations(t)
}

type tokenHandler struct{}

func (t tokenHandler) Create(token auth.Token) (*string, *auth.TokenError) {
	return nil, &auth.TokenError{
		Error:  errors.New("error"),
		Status: 500,
	}
}

func (t tokenHandler) Retrieve(token string) (*auth.Token, *auth.TokenError) {
	panic("shouldn't reach here")
}

func TestFetchUserIdTaskImpl_FailsOnTokenFailure(t *testing.T) {
	middleware := new(http_wrapper.TestMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusInternalServerError, http_wrapper.Json{
		"error": "error",
	}).Once()

	var validations []authenticated_user_validation.Validation
	fetchUserTask := func(ctx *http_wrapper.Context) (user *model.AuthenticatedUser, err error) {
		return &model.AuthenticatedUser{}, nil
	}
	fetchUserIdTask := func(credentialService service.CredentialService, user *model.AuthenticatedUser) (u uint, err error) {
		return uint(1000), nil
	}
	createTokenTask := token_task.CreateTokenTaskImpl

	postController := session.CreatePostController(
		&tokenHandler{},
		nil,
		validations,
		fetchUserTask,
		fetchUserIdTask,
		createTokenTask,
	)

	postController.Body(&http_wrapper.Context{
		Reader:     nil,
		Writer:     nil,
		Middleware: middleware,
	})

	middleware.AssertExpectations(t)
}

func TestFetchUserIdTaskImpl_SuccessReturnsToken(t *testing.T) {
	writer := new(http_wrapper.TestWriter)
	writer.On("WriteJson", http.StatusOK, http_wrapper.Json{
		"user_id": uint(1000),
		"token": "test token",
	}).Once()

	var validations []authenticated_user_validation.Validation
	fetchUserTask := func(ctx *http_wrapper.Context) (user *model.AuthenticatedUser, err error) {
		return &model.AuthenticatedUser{}, nil
	}
	fetchUserIdTask := func(credentialService service.CredentialService, user *model.AuthenticatedUser) (u uint, err error) {
		return uint(1000), nil
	}
	createTokenTask := func(userId uint, handler auth.TokenHandler) (token *string, err *auth.TokenError) {
		tokenStr := "test token"
		return &tokenStr, nil
	}

	postController := session.CreatePostController(
		nil,
		nil,
		validations,
		fetchUserTask,
		fetchUserIdTask,
		createTokenTask,
	)

	postController.Body(&http_wrapper.Context{
		Reader:     nil,
		Writer:     writer,
		Middleware: nil,
	})

	writer.AssertExpectations(t)
}