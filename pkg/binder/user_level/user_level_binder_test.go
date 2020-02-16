package user_level_test

import (
	"fmt"
	"testing"

	controllerparams "github.com/arpb2/C-3PO/pkg/controller"

	"github.com/arpb2/C-3PO/api/controller"
	userlevelbinder "github.com/arpb2/C-3PO/pkg/binder/user_level"
	"github.com/stretchr/testify/assert"
)

type MockControllerRegistrable struct {
	RegisteredControllers []controller.Controller
}

func (m *MockControllerRegistrable) Register(controller controller.Controller) {
	m.RegisteredControllers = append(m.RegisteredControllers, controller)
}

func bindControllers() *MockControllerRegistrable {
	binder := userlevelbinder.CreateBinder(nil, nil, nil)
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
	assert.NotNil(t, lookupController("GET", fmt.Sprintf("/users/:%s/levels/:%s", controllerparams.ParamUserId, controllerparams.ParamLevelId)))
	assert.NotNil(t, lookupController("PUT", fmt.Sprintf("/users/:%s/levels/:%s", controllerparams.ParamUserId, controllerparams.ParamLevelId)))
}

func TestCreateBinder_RegistersOnlyRoutes(t *testing.T) {
	registrable := bindControllers()

	assert.Equal(t, 2, len(registrable.RegisteredControllers))
}
