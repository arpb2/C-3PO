package session_binder

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/circuit_breaker"
	"github.com/arpb2/C-3PO/src/api/controller/session"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/service"
)

type binder struct{
	CircuitBreaker circuit_breaker.CircuitBreaker
	TokenHandler auth.TokenHandler
	CredentialService service.CredentialService
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	validations := []session_validation.Validation{
		session_validation.UserValidation,
		session_validation.EmailValidation,
		session_validation.PasswordValidation,
	}

	controllerRegistrable.Register(session.CreatePostController(
		b.CircuitBreaker,
		b.TokenHandler,
		b.CredentialService,
		validations,
	))
}

func CreateBinder(circuitBreaker circuit_breaker.CircuitBreaker,
	              tokenHandler auth.TokenHandler,
	              credentialService service.CredentialService) engine.ControllerBinder {
	return &binder{
		CircuitBreaker:    circuitBreaker,
		TokenHandler:      tokenHandler,
		CredentialService: credentialService,
	}
}
