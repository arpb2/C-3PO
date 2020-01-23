package hystrix

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/arpb2/C-3PO/src/api/executor"
)

func CreateExecutor() executor.Executor {
	return &Executor{}
}

var timeoutDebug = 1000000000000000

type Executor struct {}

func (cb Executor) Go(command executor.Command) chan error {
	errChan := make(chan error, 1)
	go func(errChan chan<- error) {
		err := cb.Do(command)

		if err != nil {
			errChan <- err
		}

		close(errChan)
	}(errChan)

	return errChan
}

func (cb Executor) Do(command executor.Command) error {
	if command.Prepare() {
		hystrix.ConfigureCommand(command.Name(), hystrix.CommandConfig{
			Timeout:   timeoutDebug,
		})

		return hystrix.Do(command.Name(), command.Run, command.Fallback)
	}
	return nil
}
