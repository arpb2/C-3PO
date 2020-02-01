package main

import (
	engine "github.com/arpb2/C-3PO/pkg/engine/gin"
	"github.com/arpb2/C-3PO/pkg/server"
)

func main() {
	_ = server.StartApplication(engine.New())
}
