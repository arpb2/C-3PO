package main

import (
	engine "github.com/arpb2/C-3PO/pkg/infra/engine/gin"
	"github.com/arpb2/C-3PO/pkg/infra/server"
)

func main() {
	_ = server.StartApplication(engine.New())
}
