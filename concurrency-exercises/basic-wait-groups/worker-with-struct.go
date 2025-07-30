package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID   int
	Task string
}

// PerformTask simulates a worker performing a task
func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WokerID %d started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WokerID %d finished %s\n", w.ID, w.Task)
}

func main() {
	var wg sync.WaitGroup
	// Define tasks to be performed by workers
	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{ID: i, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}

	wg.Wait()
	fmt.Println("Construction finished")
}
