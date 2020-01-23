package http_executor

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"sync"
)

func CreateHttpExecutor(exec executor.Executor) executor.HttpExecutor {
	return &httpExecutor{
		Executor: exec,
	}
}

type httpExecutor struct{
	Executor executor.Executor
}

func (e httpExecutor) BatchRun(ctx *http_wrapper.Context, commands []executor.Command) {
	doneChan := make(chan bool, 1)
	defer close(doneChan)

	var channels []<-chan error
	for _, command := range commands {
		errChan := e.Executor.Go(command)

		channels = append(channels, errChan)
	}

	fanInChannel := merge(channels...)

	for err := range fanInChannel {
		if err != nil {
			fmt.Print(err.Error())
			ctx.AbortTransactionWithError(http_wrapper.CreateInternalError())
		}
	}
}

func merge(cs ...<-chan error) <-chan error {
	var wg sync.WaitGroup

	out := make(chan error)

	send := func(c <-chan error) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}