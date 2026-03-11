package main

import (
	"context"
	"sync"
	"time"
)

type Task func(ctx context.Context) error

type TaskResult struct {
	ID       int
	Duration time.Duration
	Error    error
}

type Scheduler struct {
	maxProc int
	timeout time.Duration
}

func NewScheduler(maxProc int, timeout time.Duration) *Scheduler {
	return &Scheduler{
		maxProc: maxProc,
		timeout: timeout,
	}
}

func (s *Scheduler) Run(tasks []Task) []TaskResult {
	var wg sync.WaitGroup
	results := make([]TaskResult, len(tasks))

	sig := make(chan struct{}, s.maxProc)

	for i, task := range tasks {
		wg.Add(1)

		id := i
		t := task

		go func() {
			defer wg.Done()

			sig <- struct{}{}
			defer func() { <-sig }()

			start := time.Now()

			ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
			defer cancel()

			err := t(ctx)

			results[id] = TaskResult{
				ID:       id,
				Duration: time.Since(start),
				Error:    err,
			}
		}()
	}

	wg.Wait()
	return results
}
