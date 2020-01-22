package user_binder

import (
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_validation"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
)

type binder struct {
	Executor          executor.Executor
	AuthMiddleware http_wrapper.Handler
	UserService service.UserService
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	emptyEmailValidation := user_validation.EmptyEmailValidation
	emptyNameValidation := user_validation.EmptyNameValidation
	emptyPasswordValidation := user_validation.EmptyPasswordValidation
	emptySurnameValidation := user_validation.EmptySurnameValidation
	idProvidedValidation := user_validation.IdProvidedValidation
	securePasswordValidation := user_validation.SecurePasswordValidation

	controllerRegistrable.Register(user.CreatePostController(
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
	controllerRegistrable.Register(user.CreateGetController(
		b.Executor,
		b.AuthMiddleware,
		b.UserService))
	controllerRegistrable.Register(user.CreatePutController(
		b.Executor,
		[]user_validation.Validation{
			idProvidedValidation,
			securePasswordValidation,
		},
		b.AuthMiddleware,
		b.UserService))
	controllerRegistrable.Register(user.CreateDeleteController(
		b.Executor,
		b.AuthMiddleware,
		b.UserService))
}

func CreateBinder(exec executor.Executor, authMiddleware http_wrapper.Handler, userService service.UserService) engine.ControllerBinder {
	return &binder{
		Executor:       exec,
		AuthMiddleware: authMiddleware,
		UserService:    userService,
	}
}