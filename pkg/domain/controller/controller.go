package controller

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
)

type Controller struct {
	Method string

	Path string

	Middleware []http.Handler

	Body http.Handler
}
