package blocking

import (
	"github.com/arpb2/C-3PO/src/api/executor"
)

func CreateExecutor() executor.Executor {
	return &Executor{}
}

type Executor struct {}

func (cb Executor) Go(command executor.Command) chan error {
	errChan := make(chan error, 1)
	err := command.Run()

	if err != nil {
		err = command.Fallback(err)
	}

	if err != nil {
		errChan <- err
	}
	return errChan
}

func (cb Executor) Do(command executor.Command) error {
	err := command.Run()

	if err != nil {
		return command.Fallback(err)
	}

	return err
}
