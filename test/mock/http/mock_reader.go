package http

import "github.com/stretchr/testify/mock"

type MockReader struct {
	mock.Mock
}

func (m *MockReader) ReadBody(obj interface{}) error {
	args := m.Called(obj)

	return args.Error(0)
}

func (m *MockReader) GetMethod() string {
	args := m.Called()

	return args.String(0)
}

func (m *MockReader) GetUrl() string {
	args := m.Called()

	return args.String(0)
}

func (m *MockReader) GetParameter(key string) string {
	args := m.Called(key)

	return args.String(0)
}

func (m *MockReader) GetHeader(key string) string {
	args := m.Called(key)

	return args.String(0)
}

func (m *MockReader) GetFormData(key string) (string, bool) {
	args := m.Called(key)

	return args.String(0), args.Bool(1)
}
