package controller

import (
	"fmt"
	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	levelservice "github.com/arpb2/C-3PO/pkg/domain/service/level"
	httpcodes "net/http"
	"strconv"
)

func CreatePutController(authMiddleware http.Handler, levelService levelservice.Service) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   fmt.Sprintf("/levels/:%s", controller.ParamLevelId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(levelService),
	}
}

func CreatePutBody(levelService levelservice.Service) http.Handler {
	return func(ctx *http.Context) {
		id := ctx.GetParameter(controller.ParamLevelId)

		if id == "" {
			ctx.AbortTransactionWithError(http.CreateBadRequestError(fmt.Sprintf("'%s' empty", controller.ParamLevelId)))
			return
		}

		idUint, err := strconv.ParseUint(id, 10, 64)

		if err != nil {
			ctx.AbortTransactionWithError(http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", controller.ParamLevelId)))
			return
		}

		var level model.Level
		err = ctx.ReadBody(&level)

		if err != nil {
			ctx.AbortTransactionWithError(http.CreateBadRequestError("bad json body"))
			return
		}

		level.Id = uint(idUint)
		level, err = levelService.StoreLevel(level)

		if err != nil {
			ctx.AbortTransactionWithError(err)
			return
		}

		ctx.WriteJson(httpcodes.StatusOK, level)
	}
}
