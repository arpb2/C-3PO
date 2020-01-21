package http_wrapper

type Middleware interface {

	Next()

	Abort()

	AbortWithStatusJSON(code int, jsonObj interface{})

}
