package http_wrapper

type Reader interface {

	Url() string

	ShouldBindJSON(obj interface{}) error

	Param(key string) string

	GetHeader(key string) string

	GetPostForm(key string) (string, bool)

}
