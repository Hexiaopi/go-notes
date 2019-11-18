package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	// ErrTimeout is error when time out.
	ErrTimeout = errors.New("received timeout")
	// ErrInterrupt is error when received interrupt from operating system.
	ErrInterrupt = errors.New("received interrupt")
)

// Runner run a set of tasks within a given timeout and can be
// shut down on an operating system interrupt.
type Runner struct {
	// interrupt channel reports a signal from the operating system.
	interrupt chan os.Signal
	// complete channel reports that processing is done.
	complete chan error
	// timeout channel reports that time has run out.
	timeout <-chan time.Time
	// tasks holds a set of functions that are executed.
	tasks []func(int)
}

// New returns a new ready-to-use Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1), //必须有缓冲，否则接受不到操作系统的中断信号
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add attaches tasks to the Runner.
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// gotInterrupt verifies if the interrupt signal has been issued.
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

// run executes each registered task.
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

// Start runs all task and monitors channel events.
func (r *Runner) Start() error {
	signal.Notify(r.interrupt)
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
