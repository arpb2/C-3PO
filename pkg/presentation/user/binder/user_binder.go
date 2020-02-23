package binder

import (
	"github.com/arpb2/C-3PO/pkg/domain/engine"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	userservice "github.com/arpb2/C-3PO/pkg/domain/service/user"
	"github.com/arpb2/C-3PO/pkg/presentation/user/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"
)

type binder struct {
	Executor       pipeline.HttpPipeline
	AuthMiddleware http.Handler
	UserService    userservice.Service
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	emptyEmailValidation := validation.EmptyEmail
	emptyNameValidation := validation.EmptyName
	emptyPasswordValidation := validation.EmptyPassword
	emptySurnameValidation := validation.EmptySurname
	idProvidedValidation := validation.IdProvided
	securePasswordValidation := validation.SecurePassword

	controllerRegistrable.Register(controller.CreatePostController(
		b.Executor,
		[]validation.Validation{
			emptyEmailValidation,
			emptyNameValidation,
			emptySurnameValidation,
			emptyPasswordValidation,
			securePasswordValidation,
			idProvidedValidation,
		},
		b.UserService))
	controllerRegistrable.Register(controller.CreateGetController(
		b.Executor,
		b.AuthMiddleware,
		b.UserService))
	controllerRegistrable.Register(controller.CreatePutController(
		b.Executor,
		[]validation.Validation{
			idProvidedValidation,
			securePasswordValidation,
		},
		b.AuthMiddleware,
		b.UserService))
	controllerRegistrable.Register(controller.CreateDeleteController(
		b.Executor,
		b.AuthMiddleware,
		b.UserService))
}

func CreateBinder(exec pipeline.HttpPipeline, authMiddleware http.Handler, userService userservice.Service) engine.ControllerBinder {
	return &binder{
		Executor:       exec,
		AuthMiddleware: authMiddleware,
		UserService:    userService,
	}
}
