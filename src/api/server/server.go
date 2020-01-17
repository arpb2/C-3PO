package server

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/arpb2/C-3PO/src/api/controller/health"
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/engine"
)

func StartApplication(engine engine.ServerEngine) error {
	RegisterRoutes(engine)

	if err := engine.Run(); err != nil {
		_ = fmt.Errorf("error running server %s", err.Error())
		return err
	}
	return nil
}

type Binder func(engine.ControllerRegistrable)
var binders = []Binder{
	health.Binder,
	code.Binder,
	user.Binder,
}

func RegisterRoutes(engine engine.ServerEngine) {
	for _, binder := range binders {
		binder(engine)
	}
}
