package http_wrapper

type Reader interface {
	GetUrl() string

	GetParameter(key string) string

	GetHeader(key string) string

	GetFormData(key string) (string, bool)

	ReadBody(obj interface{}) error
}