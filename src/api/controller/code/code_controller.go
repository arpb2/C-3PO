package code

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

var GetController = controller.Controller{
	Method: "GET",
	Path:   "/users/:user_id/codes/:code_id",
	Body:   codeGet,
}

var PostController = controller.Controller{
	Method: "POST",
	Path:   "/users/:user_id/codes",
	Body:   codePost,
}

var PutController = controller.Controller{
	Method: "PUT",
	Path:   "/users/:user_id/codes/:code_id",
	Body:   codePut,
}

func codeGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": "some long code",
		"user_id": ctx.Param("user_id"),
		"code_id": ctx.Param("code_id"),
	})
}

func codePost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": "some long code",
		"user_id": ctx.Param("user_id"),
		"code_id": ctx.Param("code_id"),
	})
}

func codePut(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": "some long code",
		"user_id": ctx.Param("user_id"),
		"code_id": ctx.Param("code_id"),
	})
}
