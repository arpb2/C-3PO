package http

import "github.com/stretchr/testify/mock"

type MockMiddleware struct {
	mock.Mock
}

func (t *MockMiddleware) IsAborted() bool {
	args := t.Called()
	return args.Bool(0)
}

func (t *MockMiddleware) NextHandler() {
	_ = t.Called()
}

func (t *MockMiddleware) AbortTransaction() {
	_ = t.Called()
}

func (t *MockMiddleware) AbortTransactionWithError(err error) {
	_ = t.Called(err)
}

func (t *MockMiddleware) AbortTransactionWithStatus(code int, jsonObj interface{}) {
	_ = t.Called(code, jsonObj)
}
