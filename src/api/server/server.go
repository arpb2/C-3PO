package server

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/arpb2/C-3PO/src/api/controller/health"
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

func CreateEngine() engine.ServerEngine {
	serverEngine := engine.CreateBasicServerEngine()

	serverEngine.Register(health.GetController)

	serverEngine.Register(code.GetController)
	serverEngine.Register(code.PostController)
	serverEngine.Register(code.PutController)

	return serverEngine
}
