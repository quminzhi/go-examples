package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeout   = errors.New("timeout")
	ErrInterrupt = errors.New("interrupt")
)

type Runner struct {
	interrupt chan os.Signal // Receive interrupt signal if there is
	timeout   <-chan time.Time
	complete  chan error // return nil if no error

	tasks []func(int)
}

func NewRunner(timeout time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1), // reserve enough space for signals, refer to signal.Notify()
		timeout:   time.After(timeout),
		complete:  make(chan error), // return complete status
		tasks:     make([]func(int), 0),
	}
}

func (r *Runner) run() error {
	for tid, task := range r.tasks {
		// Check if interrupt received
		select {
		case <-r.interrupt:
			signal.Stop(r.interrupt)
			return ErrInterrupt
		default:
			task(tid)
		}
	}
	return nil
}

func (r *Runner) AddTasks(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}
