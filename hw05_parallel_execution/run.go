package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tasksCh := make(chan Task)
	var tasksErrors int32

	// if m <= 0, ignore errors
	if m <= 0 {
		m = len(tasks) + 1
	}

	var wg sync.WaitGroup
	wg.Add(n)

	worker := func() {
		defer wg.Done()
		for t := range tasksCh {
			err := t()
			if err != nil {
				atomic.AddInt32(&tasksErrors, 1)
			}
		}
	}

	// Start n goroutines
	for i := 0; i < n; i++ {
		go worker()
	}

	// Put tasks into goroutines, if there errors - break
	for i := 0; i < len(tasks); i++ {
		if atomic.LoadInt32(&tasksErrors) >= int32(m) {
			break
		}
		tasksCh <- tasks[i]
	}
	close(tasksCh)
	wg.Wait()

	if tasksErrors >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
