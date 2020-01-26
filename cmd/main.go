package main

import (
	engine "github.com/arpb2/C-3PO/internal/engine/gin"
	"github.com/arpb2/C-3PO/internal/server"
)

func main() {
	_ = server.StartApplication(engine.New())
}
