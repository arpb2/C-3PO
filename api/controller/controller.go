package controller

import (
	"github.com/arpb2/C-3PO/api/http"
)

type Controller struct {
	Method string

	Path string

	Middleware []http.Handler

	Body http.Handler
}
