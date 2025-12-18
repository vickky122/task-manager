package manager

import (
	"container/heap"
	"log"
	"sync"
	"time"

	"task-manager/model"
	"task-manager/queue"
)

type TaskManager struct {
	queue  *queue.TaskQueue
	mu     sync.Mutex
	cond   *sync.Cond
	taskID int
}

func NewTaskManager(workers int) *TaskManager {
	tm := &TaskManager{
		queue: queue.InitQueue(),
	}
	tm.cond = sync.NewCond(&tm.mu)

	for i := 0; i < workers; i++ {
		go tm.worker(i)
	}
	return tm
}

func (tm *TaskManager) AddTask(name string, priority int) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.taskID++
	task := &model.Task{
		ID:       tm.taskID,
		Name:     name,
		Priority: priority,
	}

	heap.Push(tm.queue, task)
	tm.cond.Signal()
}

func (tm *TaskManager) worker(id int) {
	for {
		tm.mu.Lock()
		for tm.queue.Len() == 0 {
			tm.cond.Wait()
		}
		task := heap.Pop(tm.queue).(*model.Task)
		tm.mu.Unlock()

		log.Printf("Worker %d executing task %d (%s) priority=%d",
			id, task.ID, task.Name, task.Priority)

		time.Sleep(1 * time.Second)
	}
}
