package binder

import (
	"github.com/arpb2/C-3PO/pkg/domain/auth"
	"github.com/arpb2/C-3PO/pkg/domain/engine"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	credentialservice "github.com/arpb2/C-3PO/pkg/domain/service/credential"
	"github.com/arpb2/C-3PO/pkg/presentation/session/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"
)

type binder struct {
	Executor          pipeline.HttpPipeline
	TokenHandler      auth.TokenHandler
	CredentialService credentialservice.Service
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	validations := []validation.Validation{
		validation.EmptyUser,
		validation.EmptyEmail,
		validation.EmptyPassword,
	}

	controllerRegistrable.Register(controller.CreatePostController(
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
