package http_wrapper

import "github.com/stretchr/testify/mock"

type TestReader struct{
	mock.Mock
}

func (m TestReader) ShouldBindJSON(obj interface{}) error {
	args := m.Called(obj)

	return args.Error(0)
}

func (m TestReader) Url() string {
	args := m.Called()

	return args.String(0)
}

func (m TestReader) Param(key string) string {
	args := m.Called(key)

	return args.String(0)
}

func (m TestReader) GetHeader(key string) string {
	args := m.Called(key)

	return args.String(0)
}

func (m TestReader) GetPostForm(key string) (string, bool) {
	args := m.Called(key)

	return args.String(0), args.Bool(1)
}
