package user

import (
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
)

type binder struct {
	AuthMiddleware http_wrapper.Handler
	UserService service.UserService
}

func (b binder) BindControllers(controllerRegistrable engine.ControllerRegistrable) {
	controllerRegistrable.Register(CreatePostController(b.UserService))

	controllerRegistrable.Register(CreateGetController(b.AuthMiddleware, b.UserService))
	controllerRegistrable.Register(CreatePutController(b.AuthMiddleware, b.UserService))
	controllerRegistrable.Register(CreateDeleteController(b.AuthMiddleware, b.UserService))
}

func CreateBinder(authMiddleware http_wrapper.Handler, userService service.UserService) engine.ControllerBinder {
	return &binder{
		AuthMiddleware: authMiddleware,
		UserService:    userService,
	}
}