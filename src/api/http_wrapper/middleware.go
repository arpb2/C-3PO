package http_wrapper

type Middleware interface {

	NextHandler()

	AbortTransaction()

	AbortTransactionWithStatus(code int, jsonObj interface{})

}
