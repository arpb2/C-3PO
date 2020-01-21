package code

import (
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/teacher_auth"
	"github.com/arpb2/C-3PO/src/api/service/code_service"
	"github.com/arpb2/C-3PO/src/api/service/teacher_service"
	"net/http"
	"strconv"
)

func Binder(handler engine.ControllerRegistrable) {
	authMiddleware := teacher_auth.CreateMiddleware(jwt.CreateTokenHandler(), teacher_service.GetService())
	codeService := code_service.GetService()

	handler.Register(CreateGetController(authMiddleware, codeService))
	handler.Register(CreatePostController(authMiddleware, codeService))
	handler.Register(CreatePutController(authMiddleware, codeService))
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

func FetchCodeId(ctx *http_wrapper.Context) (uint, bool) {
	codeId := ctx.GetParameter("code_id")

	if codeId == "" {
		controller.Halt(ctx, http.StatusBadRequest, "'code_id' empty")
		return 0, true
	}

	codeIdUint, err := strconv.ParseUint(codeId, 10, 64)

	if err != nil {
		controller.Halt(ctx, http.StatusBadRequest, "'code_id' malformed. Expecting a positive number.")
		return 0, true
	}

	return uint(codeIdUint), false
}

func FetchCode(ctx *http_wrapper.Context) (*string, bool) {
	code, exists := ctx.GetFormData("code")

	if !exists {
		controller.Halt(ctx, http.StatusBadRequest, "'code' part not found")
		return nil, true
	}

	return &code, false
}