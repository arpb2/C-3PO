package code_binder_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/controller"
	code_binder "github.com/arpb2/C-3PO/pkg/binder/code"
	"github.com/stretchr/testify/assert"
)

type MockControllerRegistrable struct {
	RegisteredControllers []controller.Controller
}

func (m *MockControllerRegistrable) Register(controller controller.Controller) {
	m.RegisteredControllers = append(m.RegisteredControllers, controller)
}

func bindControllers() *MockControllerRegistrable {
	binder := code_binder.CreateBinder(nil, nil, nil)
	registrable := &MockControllerRegistrable{}

	binder.BindControllers(registrable)

	return registrable
}

func lookupController(method, path string) *controller.Controller {
	registrable := bindControllers()
	var expectedController *controller.Controller

	for _, registeredController := range registrable.RegisteredControllers {
		if registeredController.Method == method && registeredController.Path == path {
			expectedController = &registeredController
		}
	}

	return expectedController
}

func TestCreateBinder_RegistersRoutes(t *testing.T) {
	assert.NotNil(t, lookupController("GET", "/users/:user_id/codes/:code_id"))
	assert.NotNil(t, lookupController("POST", "/users/:user_id/codes"))
	assert.NotNil(t, lookupController("PUT", "/users/:user_id/codes/:code_id"))
}

func TestCreateBinder_RegistersOnlyRoutes(t *testing.T) {
	registrable := bindControllers()

	assert.Equal(t, 3, len(registrable.RegisteredControllers))
}
