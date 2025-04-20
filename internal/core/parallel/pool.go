// Package parallel provides utilities for parallel processing and concurrency management.
package parallel

import (
	"context"
	"sync"
)

// WorkerPool represents a pool of workers that can execute tasks concurrently.
// It manages a fixed number of goroutines (workers) that process tasks from a queue.
type WorkerPool struct {
	taskQueue chan func()    // Channel for queuing tasks to be executed by workers
	workers   int            // Number of worker goroutines to spawn
	wg        sync.WaitGroup // WaitGroup to track completion of all workers
	once      sync.Once      // Ensures the taskQueue is closed only once
}

// NewWorkerPool creates a new worker pool with the specified number of workers and buffer size.
// It initializes the task queue and starts the worker goroutines.
//
// Parameters:
//   - ctx: Context for controlling the lifecycle of the worker pool
//   - workers: Number of worker goroutines to create
//   - buffer: Size of the task queue buffer
//
// Returns:
//   - A pointer to the initialized WorkerPool
func NewWorkerPool(ctx context.Context, workers, buffer int) *WorkerPool {
	p := &WorkerPool{
		taskQueue: make(chan func(), buffer),
		workers:   workers,
	}
	go p.dispatch(ctx)
	return p
}

// AddTask adds a new task to the worker pool's queue.
// The task will be executed by one of the worker goroutines when it becomes available.
//
// Parameters:
//   - task: A function with no parameters and no return value to be executed by a worker
func (p *WorkerPool) AddTask(task func()) {
	p.taskQueue <- task
}

// Wait closes the task queue and waits for all worker goroutines to complete their tasks.
// This method should be called after all tasks have been added to the pool.
// It ensures that the task queue is closed only once using sync.Once.
func (p *WorkerPool) Wait() {
	p.once.Do(func() {
		close(p.taskQueue)
	})
	p.wg.Wait()
}

// dispatch starts the specified number of worker goroutines.
// Each worker processes tasks from the task queue until the queue is closed or the context is done.
//
// Parameters:
//   - ctx: Context for controlling the lifecycle of the workers
func (p *WorkerPool) dispatch(ctx context.Context) {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go Worker(ctx, p.taskQueue, &p.wg)
	}
}
