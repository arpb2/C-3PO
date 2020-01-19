package user

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

var GetController = controller.Controller{
	Method: "GET",
	Path:   "/users/:user_id",
	Body:   userGet,
}

func userGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{})
}