package user_binder

import (
	"github.com/arpb2/C-3PO/api/engine"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/pipeline"
	user_service "github.com/arpb2/C-3PO/api/service/user"
	user_controller "github.com/arpb2/C-3PO/pkg/controller/user"
	user_validation "github.com/arpb2/C-3PO/pkg/validation/user"
)

type binder struct {
	Executor       pipeline.HttpPipeline
	AuthMiddleware http_wrapper.Handler
	UserService    user_service.Service
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	emptyEmailValidation := user_validation.EmptyEmail
	emptyNameValidation := user_validation.EmptyName
	emptyPasswordValidation := user_validation.EmptyPassword
	emptySurnameValidation := user_validation.EmptySurname
	idProvidedValidation := user_validation.IdProvided
	securePasswordValidation := user_validation.SecurePassword

	controllerRegistrable.Register(user_controller.CreatePostController(
		b.Executor,
		[]user_validation.Validation{
			emptyEmailValidation,
			emptyNameValidation,
			emptySurnameValidation,
			emptyPasswordValidation,
			securePasswordValidation,
			idProvidedValidation,
		},
		b.UserService))
	controllerRegistrable.Register(user_controller.CreateGetController(
		b.Executor,
		b.AuthMiddleware,
		b.UserService))
	controllerRegistrable.Register(user_controller.CreatePutController(
		b.Executor,
		[]user_validation.Validation{
			idProvidedValidation,
			securePasswordValidation,
		},
		b.AuthMiddleware,
		b.UserService))
	controllerRegistrable.Register(user_controller.CreateDeleteController(
		b.Executor,
		b.AuthMiddleware,
		b.UserService))
}

func CreateBinder(exec pipeline.HttpPipeline, authMiddleware http_wrapper.Handler, userService user_service.Service) engine.ControllerBinder {
	return &binder{
		Executor:       exec,
		AuthMiddleware: authMiddleware,
		UserService:    userService,
	}
}
