package main

import (
	"sync/atomic"
	"task-manager/model"
	"task-manager/scheduler"
	"task-manager/worker"
)

var idCounter int64

func main() {
	s := scheduler.NewScheduler()

	for i := 1; i <= 3; i++ {
		worker.StartWorker(i, s.Tasks())
	}

	s.Submit(&model.Task{
		ID:       atomic.AddInt64(&idCounter, 1),
		Name:     "Low priority",
		Priority: 1,
	})

	s.Submit(&model.Task{
		ID:       atomic.AddInt64(&idCounter, 1),
		Name:     "High priority",
		Priority: 10,
	})

	select {}
}
