package queue

import (
	"container/heap"
	"task-manager/model"
)

type TaskQueue []*model.Task

func (pq TaskQueue) Len() int {
	return len(pq)
}

func (pq TaskQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq TaskQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *TaskQueue) Push(x interface{}) {
	task := x.(*model.Task)
	*pq = append(*pq, task)
}

func (pq *TaskQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	*pq = old[:n-1]
	return task
}

func InitQueue() *TaskQueue {
	pq := &TaskQueue{}
	heap.Init(pq)
	return pq
}
