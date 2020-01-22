package hystrix

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/arpb2/C-3PO/src/api/executor"
)

func CreateExecutor() executor.Executor {
	return &Executor{}
}

type Executor struct {}

func (cb Executor) Go(command executor.Command) chan error {
	return hystrix.Go(command.Name(), command.Run, command.Fallback)
}

func (cb Executor) Do(command executor.Command) error {
	return hystrix.Do(command.Name(), command.Run, command.Fallback)
}
