package http_wrapper

import "github.com/stretchr/testify/mock"

type MockWriter struct {
	mock.Mock
}

func (t MockWriter) WriteJson(code int, obj interface{}) {
	_ = t.Called(code, obj)
}

func (t MockWriter) WriteString(code int, format string, values ...interface{}) {
	_ = t.Called(code, format, values)
}

func (t MockWriter) WriteStatus(code int) {
	_ = t.Called(code)
}
