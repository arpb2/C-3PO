package code

import (
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Service service.CodeService

func Binder(handler engine.ControllerHandler) {
	handler.Register(GetController)
	handler.Register(PostController)
	handler.Register(PutController)
}

func FetchUserId(ctx *gin.Context) (string, bool) {
	userId := ctx.Param("user_id")

	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "'user_id' empty",
		})
		return userId, true
	}

	return userId, false
}

func FetchCodeId(ctx *gin.Context) (string, bool) {
	codeId := ctx.Param("code_id")

	if codeId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "'code_id' empty",
		})
		return codeId, true
	}

	return codeId, false
}

func FetchCode(ctx *gin.Context) (*string, bool) {
	code, exists := ctx.GetPostForm("code")

	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "'code' part not found",
		})
		return nil, true
	}

	return &code, false
}