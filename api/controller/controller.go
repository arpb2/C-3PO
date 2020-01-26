package controller

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
)

type Controller struct {
	Method string

	Path string

	Middleware []http_wrapper.Handler

	Body http_wrapper.Handler
}
