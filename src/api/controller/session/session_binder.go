package session

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/arpb2/C-3PO/src/api/task/authenticated_user_task"
	"github.com/arpb2/C-3PO/src/api/task/token_task"
	"github.com/arpb2/C-3PO/src/api/validation/authenticated_user_validation"
)

type binder struct{
	TokenHandler auth.TokenHandler
	CredentialService service.CredentialService
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	validations := []authenticated_user_validation.Validation{
		authenticated_user_validation.UserValidation,
		authenticated_user_validation.EmailValidation,
		authenticated_user_validation.PasswordValidation,
	}
	fetchUserTask := authenticated_user_task.FetchUserTaskImpl
	fetchUserIdTask := authenticated_user_task.FetchUserIdTaskImpl
	createTokenTask := token_task.CreateTokenTaskImpl

	controllerRegistrable.Register(CreatePostController(
		b.TokenHandler,
		b.CredentialService,
		validations,
		fetchUserTask,
		fetchUserIdTask,
		createTokenTask,
	))
}

func CreateBinder(tokenHandler auth.TokenHandler, credentialService service.CredentialService) engine.ControllerBinder {
	return &binder{
		TokenHandler:      tokenHandler,
		CredentialService: credentialService,
	}
}
