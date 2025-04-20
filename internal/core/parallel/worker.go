// Package parallel provides utilities for parallel processing and concurrency management.
package parallel

import (
	"context"
	"sync"
)

// Worker is a function that processes tasks from the task queue.
// It continues processing tasks until either the task queue is closed or the context is done.
//
// Parameters:
//   - ctx: Context for controlling the lifecycle of the worker
//   - taskQueue: Channel from which tasks are received
//   - wg: WaitGroup to signal when the worker has completed
func Worker(ctx context.Context, taskQueue chan func(), wg *sync.WaitGroup) {
	defer wg.Done()
Loop:
	for {
		select {
		case task, ok := <-taskQueue:
			if !ok {
				break Loop
			}
			task()
		case <-ctx.Done():
			break Loop
		}
	}
}
