package controller

import "github.com/gin-gonic/gin"

type Controller struct {
	Method string

	Path string

	Middleware []gin.HandlerFunc

	Body gin.HandlerFunc
}
