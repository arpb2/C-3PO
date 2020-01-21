package user

import (
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
)

func Binder(handler engine.ControllerRegistrable) {
	authMiddleware := single_auth.CreateMiddleware(jwt.CreateTokenHandler())

	handler.Register(CreateGetController(authMiddleware))
	handler.Register(CreatePostController(authMiddleware))
	handler.Register(CreatePutController(authMiddleware))
	handler.Register(CreateDeleteController(authMiddleware))
}