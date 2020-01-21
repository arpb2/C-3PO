package http_wrapper

type Json map[string]interface{}

type Writer interface {

	JSON(code int, obj interface{})

	String(code int, format string, values ...interface{})

}
