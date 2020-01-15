package server

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/arpb2/C-3PO/src/api/controller/health"
	"github.com/arpb2/C-3PO/src/api/controller/user"
	"github.com/arpb2/C-3PO/src/api/engine"
)

var Engine = CreateEngine()

func StartApplication() error {
	if err := Engine.Run(); err != nil {
		_ = fmt.Errorf("error running server %s", err.Error())
		return err
	}
	return nil
}

type Binder func(engine.ControllerHandler)
var binders = []Binder{
	health.Binder,
	code.Binder,
	user.Binder,
}

func CreateEngine() engine.ServerEngine {
	serverEngine := engine.CreateBasicServerEngine()

	for _, binder := range binders {
		binder(serverEngine)
	}

	return serverEngine
}
