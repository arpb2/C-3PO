package code

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Binder(handler engine.ControllerRegistrable) {
	handler.Register(CreateGetController())
	handler.Register(CreatePostController())
	handler.Register(CreatePutController())
}

func FetchUserId(ctx *gin.Context) (string, bool) {
	userId := ctx.Param("user_id")

	if userId == "" {
		controller.Halt(ctx, http.StatusBadRequest, "'user_id' empty")
		return userId, true
	}

	return userId, false
}

func FetchCodeId(ctx *gin.Context) (string, bool) {
	codeId := ctx.Param("code_id")

	if codeId == "" {
		controller.Halt(ctx, http.StatusBadRequest, "'code_id' empty")
		return codeId, true
	}

	return codeId, false
}

func FetchCode(ctx *gin.Context) (*string, bool) {
	code, exists := ctx.GetPostForm("code")

	if !exists {
		controller.Halt(ctx, http.StatusBadRequest, "'code' part not found")
		return nil, true
	}

	return &code, false
}