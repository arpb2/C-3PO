package executor

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
)

type HttpExecutor interface {

	BatchRun(ctx *http_wrapper.Context, commands []Command)

}