package main

import (
	engine "github.com/arpb2/C-3PO/src/api/engine/c3po"
	"github.com/arpb2/C-3PO/src/api/server"
)

func main() {
	_ = server.StartApplication(engine.New())
}
