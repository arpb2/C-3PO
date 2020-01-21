package user

import (
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service/user_service"
	"net/http"
	"strconv"
)

func Binder(handler engine.ControllerRegistrable) {
	authMiddleware := single_auth.CreateMiddleware(jwt.CreateTokenHandler())
	userService := user_service.GetService()

	handler.Register(CreateGetController(authMiddleware, userService))
	handler.Register(CreatePostController(authMiddleware, userService))
	handler.Register(CreatePutController(authMiddleware, userService))
	handler.Register(CreateDeleteController(authMiddleware, userService))
}

func FetchUserId(ctx *http_wrapper.Context) (uint, bool) {
	userId := ctx.GetParameter("user_id")

	if userId == "" {
		controller.Halt(ctx, http.StatusBadRequest, "'user_id' empty")
		return 0, true
	}

	userIdUint, err := strconv.ParseUint(userId, 10, 64)

	if err != nil {
		controller.Halt(ctx, http.StatusBadRequest, "'user_id' malformed. Expecting a positive number.")
		return 0, true
	}

	return uint(userIdUint), false
}

func FetchAuthenticatedUser(ctx *http_wrapper.Context) (*model.AuthenticatedUser, bool) {
	var authenticatedUser model.AuthenticatedUser

	if err := ctx.ReadBody(&authenticatedUser); err != nil {
		controller.Halt(ctx, http.StatusBadRequest, "bad 'user' body")
		return nil, true
	}

	return &authenticatedUser, false
}