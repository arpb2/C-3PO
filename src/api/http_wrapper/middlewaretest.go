package http_wrapper

import "github.com/stretchr/testify/mock"

type TestMiddleware struct{
	mock.Mock
}

func (t TestMiddleware) IsAborted() bool {
	args := t.Called()
	return args.Bool(0)
}

func (t TestMiddleware) NextHandler() {
	_ = t.Called()
}

func (t TestMiddleware) AbortTransaction() {
	_ = t.Called()
}

func (t TestMiddleware) AbortTransactionWithStatus(code int, jsonObj interface{}) {
	_ = t.Called(code, jsonObj)
}
