package user

import (
	"github.com/arpb2/C-3PO/api/engine"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	userservice "github.com/arpb2/C-3PO/api/service/user"
	usercontroller "github.com/arpb2/C-3PO/pkg/controller/user"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
)

type binder struct {
	Executor       pipeline.HttpPipeline
	AuthMiddleware http.Handler
	UserService    userservice.Service
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	emptyEmailValidation := uservalidation.EmptyEmail
	emptyNameValidation := uservalidation.EmptyName
	emptyPasswordValidation := uservalidation.EmptyPassword
	emptySurnameValidation := uservalidation.EmptySurname
	idProvidedValidation := uservalidation.IdProvided
	securePasswordValidation := uservalidation.SecurePassword

	controllerRegistrable.Register(usercontroller.CreatePostController(
		b.Executor,
		[]uservalidation.Validation{
			emptyEmailValidation,
			emptyNameValidation,
			emptySurnameValidation,
			emptyPasswordValidation,
			securePasswordValidation,
			idProvidedValidation,
		},
		b.UserService))
	controllerRegistrable.Register(usercontroller.CreateGetController(
		b.Executor,
		b.AuthMiddleware,
		b.UserService))
	controllerRegistrable.Register(usercontroller.CreatePutController(
		b.Executor,
		[]uservalidation.Validation{
			idProvidedValidation,
			securePasswordValidation,
		},
		b.AuthMiddleware,
		b.UserService))
	controllerRegistrable.Register(usercontroller.CreateDeleteController(
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
