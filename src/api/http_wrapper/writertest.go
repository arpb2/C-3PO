package http_wrapper

import "github.com/stretchr/testify/mock"

type TestWriter struct{
	mock.Mock
}

func (t TestWriter) WriteJson(code int, obj interface{}) {
	_ = t.Called(code, obj)
}

func (t TestWriter) WriteString(code int, format string, values ...interface{}) {
	_ = t.Called(code, format, values)
}

func (t TestWriter) WriteStatus(code int) {
	_ = t.Called(code)
}
