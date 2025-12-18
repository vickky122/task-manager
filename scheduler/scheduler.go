package scheduler

import (
	"container/heap"
	"task-manager/model"
)

type Scheduler struct {
	input  chan *model.Task
	output chan *model.Task
	pq     PriorityQueue
}

func NewScheduler() *Scheduler {
	s := &Scheduler{
		input:  make(chan *model.Task),
		output: make(chan *model.Task),
		pq:     PriorityQueue{},
	}
	heap.Init(&s.pq)
	go s.run()
	return s
}

func (s *Scheduler) run() {
	for {
		select {
		case task := <-s.input:
			heap.Push(&s.pq, task)

		case s.output <- s.pop():
			// send highest priority task to worker
		}
	}
}

func (s *Scheduler) pop() *model.Task {
	if len(s.pq) == 0 {
		return nil
	}
	return heap.Pop(&s.pq).(*model.Task)
}

func (s *Scheduler) Submit(task *model.Task) {
	s.input <- task
}

func (s *Scheduler) Tasks() <-chan *model.Task {
	return s.output
}
