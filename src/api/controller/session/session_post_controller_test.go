package session_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/session"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_task"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
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

	reader := new(http_wrapper.TestReader)
	reader.On("ReadBody", mock.MatchedBy(func(obj interface{}) bool {
		return true
	})).Return(nil).Once()

	middleware := new(http_wrapper.TestMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusBadRequest, http_wrapper.Json{
		"error": err.Error(),
	}).Once()

	validations := []session_validation.Validation{
		func(user *model.AuthenticatedUser) error {
			return nil
		},
		func(user *model.AuthenticatedUser) error {
			return err
		},
	}
	fetchUserTask := session_task.CreateFetchUserTask()

	postController := session.CreatePostController(
		nil,
		nil,
		validations,
		fetchUserTask,
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

	var validations []session_validation.Validation
	fetchUserTask := func(ctx *http_wrapper.Context) (user *model.AuthenticatedUser, err error) {
		return &model.AuthenticatedUser{
			User: &model.User{
				Email: "test@email.com",
			},
			Password: "testpassword",
		}, nil
	}

	postController := session.CreatePostController(
		nil,
		service,
		validations,
		fetchUserTask,
	)

	postController.Body(&http_wrapper.Context{
		Reader:     nil,
		Writer:     nil,
		Middleware: middleware,
	})

	middleware.AssertExpectations(t)
	service.AssertExpectations(t)
}

type tokenHandler struct{
	mock.Mock
}

func (t tokenHandler) Create(token *auth.Token) (tokenStr string, err *auth.TokenError) {
	args := t.Called(token)

	tokenStr = args.String(0)

	errParam := args.Get(1)
	if errParam != nil {
		err = errParam.(*auth.TokenError)
	}

	return
}

func (t tokenHandler) Retrieve(token string) (*auth.Token, *auth.TokenError) {
	panic("shouldn't reach here")
}

func TestFetchUserIdTaskImpl_FailsOnTokenFailure(t *testing.T) {
	middleware := new(http_wrapper.TestMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusInternalServerError, http_wrapper.Json{
		"error": "error",
	}).Once()

	var validations []session_validation.Validation
	fetchUserTask := func(ctx *http_wrapper.Context) (user *model.AuthenticatedUser, err error) {
		return &model.AuthenticatedUser{
			User: &model.User{
				Id: 1000,
			},
		}, nil
	}

	credentialService := new(credentialService)
	credentialService.On("Retrieve", "", "").Return(uint(1000), nil)

	tokenHandler := new(tokenHandler)
	tokenHandler.On("Create", mock.MatchedBy(func(tkn *auth.Token) bool {
		return tkn.UserId == uint(1000)
	})).Return("", &auth.TokenError{
		Error:  errors.New("error"),
		Status: 500,
	})

	postController := session.CreatePostController(
		tokenHandler,
		credentialService,
		validations,
		fetchUserTask,
	)

	postController.Body(&http_wrapper.Context{
		Reader:     nil,
		Writer:     nil,
		Middleware: middleware,
	})

	middleware.AssertExpectations(t)
	credentialService.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func TestFetchUserIdTaskImpl_SuccessReturnsToken(t *testing.T) {
	writer := new(http_wrapper.TestWriter)
	writer.On("WriteJson", http.StatusOK, http_wrapper.Json{
		"user_id": uint(1000),
		"token": "test token",
	}).Once()

	var validations []session_validation.Validation
	fetchUserTask := func(ctx *http_wrapper.Context) (user *model.AuthenticatedUser, err error) {
		return &model.AuthenticatedUser{
			User: &model.User{
				Email: "test@email.com",
			},
			Password: "test password",
		}, nil
	}

	credentialService := new(credentialService)
	credentialService.On("Retrieve", "test@email.com", "test password").Return(uint(1000), nil)

	tokenHandler := new(tokenHandler)
	tokenHandler.On("Create", mock.MatchedBy(func(tkn *auth.Token) bool {
		return tkn.UserId == uint(1000)
	})).Return("test token", nil)

	postController := session.CreatePostController(
		tokenHandler,
		credentialService,
		validations,
		fetchUserTask,
	)

	postController.Body(&http_wrapper.Context{
		Reader:     nil,
		Writer:     writer,
		Middleware: nil,
	})

	writer.AssertExpectations(t)
	credentialService.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}