package user

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

var PostController = controller.Controller{
	Method: "POST",
	Path:   "/users",
	Body:   userPost,
}

func userPost(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}
