package controller

import (
	"fmt"
	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	levelservice "github.com/arpb2/C-3PO/pkg/domain/service/level"
	httpcodes "net/http"
	"strconv"
)

func CreateGetController(levelService levelservice.Service) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   fmt.Sprintf("/levels/:%s", controller.ParamLevelId),
		Body:   CreateGetBody(levelService),
	}
}

func CreateGetBody(levelService levelservice.Service) http.Handler {
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

		level, err := levelService.GetLevel(uint(idUint))

		if err != nil {
			ctx.AbortTransactionWithError(err)
			return
		}

		ctx.WriteJson(httpcodes.StatusOK, level)
	}
}
