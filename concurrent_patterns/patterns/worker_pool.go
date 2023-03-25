package patterns

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID       int
	WorkerID int
}

func worker(workerID int, wg *sync.WaitGroup, tasks <-chan Task, results chan<- Task) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("I am worker %v, work on task with id %v\n", workerID, task.ID)
		time.Sleep(time.Second)
		results <- Task{
			ID:       task.ID,
			WorkerID: workerID,
		}
	}
}

func workerPool() {
	tasks := make(chan Task, 3)
	for i := 0; i < 3; i++ {
		tasks <- Task{
			ID:       i,
			WorkerID: -1,
		}
	}
	close(tasks)

	var wg sync.WaitGroup
	results := make(chan Task, 3)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go worker(i, &wg, tasks, results)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result.ID, result.WorkerID)
	}
}
