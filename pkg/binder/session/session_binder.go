package session_binder

import (
	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/engine"
	"github.com/arpb2/C-3PO/api/pipeline"
	credential_service "github.com/arpb2/C-3PO/api/service/credential"
	session_controller "github.com/arpb2/C-3PO/pkg/controller/session"
	user_validation "github.com/arpb2/C-3PO/pkg/validation/user"
)

type binder struct {
	Executor          pipeline.HttpPipeline
	TokenHandler      auth.TokenHandler
	CredentialService credential_service.Service
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	validations := []user_validation.Validation{
		user_validation.EmptyUser,
		user_validation.EmptyEmail,
		user_validation.EmptyPassword,
	}

	controllerRegistrable.Register(session_controller.CreatePostController(
		b.Executor,
		b.TokenHandler,
		b.CredentialService,
		validations,
	))
}

func CreateBinder(executor pipeline.HttpPipeline,
	tokenHandler auth.TokenHandler,
	credentialService credential_service.Service) engine.ControllerBinder {
	return &binder{
		Executor:          executor,
		TokenHandler:      tokenHandler,
		CredentialService: credentialService,
	}
}
