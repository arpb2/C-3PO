package session_binder

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/controller/session"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/service"
)

type binder struct{
	Executor          executor.Executor
	TokenHandler      auth.TokenHandler
	CredentialService service.CredentialService
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	validations := []session_validation.Validation{
		session_validation.UserValidation,
		session_validation.EmailValidation,
		session_validation.PasswordValidation,
	}

	controllerRegistrable.Register(session.CreatePostController(
		b.Executor,
		b.TokenHandler,
		b.CredentialService,
		validations,
	))
}

func CreateBinder(executor executor.Executor,
	              tokenHandler auth.TokenHandler,
	              credentialService service.CredentialService) engine.ControllerBinder {
	return &binder{
		Executor:          executor,
		TokenHandler:      tokenHandler,
		CredentialService: credentialService,
	}
}
