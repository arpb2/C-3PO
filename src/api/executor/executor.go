package executor

import "sync"

type Command interface {

	Name() string

	Run() error

	Fallback(err error) error

}

type Executor interface {

	Go(command Command) chan error

	Do(command Command) error

}

func Merge(cs ...<-chan error) <-chan error {
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
