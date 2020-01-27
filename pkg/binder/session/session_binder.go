package session

import (
	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/engine"
	"github.com/arpb2/C-3PO/api/pipeline"
	credentialservice "github.com/arpb2/C-3PO/api/service/credential"
	sessioncontroller "github.com/arpb2/C-3PO/pkg/controller/session"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
)

type binder struct {
	Executor          pipeline.HttpPipeline
	TokenHandler      auth.TokenHandler
	CredentialService credentialservice.Service
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	validations := []uservalidation.Validation{
		uservalidation.EmptyUser,
		uservalidation.EmptyEmail,
		uservalidation.EmptyPassword,
	}

	controllerRegistrable.Register(sessioncontroller.CreatePostController(
		b.Executor,
		b.TokenHandler,
		b.CredentialService,
		validations,
	))
}

func CreateBinder(executor pipeline.HttpPipeline,
	tokenHandler auth.TokenHandler,
	credentialService credentialservice.Service) engine.ControllerBinder {
	return &binder{
		Executor:          executor,
		TokenHandler:      tokenHandler,
		CredentialService: credentialService,
	}
}
