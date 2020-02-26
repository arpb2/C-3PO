package server

import (
	"fmt"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/controller"
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/engine"
)

func StartApplication(engine engine.ServerEngine, controllers []controller.Controller) error {
	for _, c := range controllers {
		engine.Register(c)
	}

	if err := engine.Run(); err != nil {
		_ = fmt.Errorf("error running server %s", err.Error())
		return err
	}
	return nil
}
