package http_wrapper

type Handler func(ctx *Context)

type Context struct {
	Reader
	Writer
	Middleware
}
