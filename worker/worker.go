package worker

import (
	"fmt"
	"task-manager/model"
)

func StartWorker(id int, tasks <-chan *model.Task) {
	go func() {
		for task := range tasks {
			if task == nil {
				continue
			}
			fmt.Printf("Worker %d executing task %d (%s)\n",
				id, task.ID, task.Name)
		}
	}()
}
