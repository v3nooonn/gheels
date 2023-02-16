package messager

import (
	"context"
)

type (
	// Executor abstraction of execution
	Executor interface {
		Execute(ctx context.Context) error
	}
	// Executors abstraction of composition
	Executors interface {
		Executor
		Add(ntfs ...Executor)
		List() []Executor
	}
)

// Notifier the implementation of Executors
type Notifier struct {
	executors []Executor
}

func (n *Notifier) Execute(ctx context.Context) error {
	for _, executor := range n.executors {
		if err := executor.Execute(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (n *Notifier) Add(executors ...Executor) {
	n.executors = append(n.executors, executors...)
}

func (n *Notifier) List() []Executor {
	return n.executors
}
