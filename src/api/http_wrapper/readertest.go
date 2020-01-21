package http_wrapper

import "github.com/stretchr/testify/mock"

type TestReader struct{
	mock.Mock
}

func (m TestReader) ReadBody(obj interface{}) error {
	args := m.Called(obj)

	return args.Error(0)
}

func (m TestReader) GetUrl() string {
	args := m.Called()

	return args.String(0)
}

func (m TestReader) GetParameter(key string) string {
	args := m.Called(key)

	return args.String(0)
}

func (m TestReader) GetHeader(key string) string {
	args := m.Called(key)

	return args.String(0)
}

func (m TestReader) GetFormData(key string) (string, bool) {
	args := m.Called(key)

	return args.String(0), args.Bool(1)
}
