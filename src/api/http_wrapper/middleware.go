package http_wrapper

type Middleware interface {

	Next()

	IsAborted() bool

	Abort()

	AbortWithStatusJSON(code int, jsonObj interface{})

}
