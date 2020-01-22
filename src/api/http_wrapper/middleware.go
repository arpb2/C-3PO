package http_wrapper

type Middleware interface {

	NextHandler()

	IsAborted() bool

	AbortTransaction()

	AbortTransactionWithStatus(code int, jsonObj interface{})

}
