package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		// Simulate some work
		time.Sleep(time.Second)
		results <- task * 2
	}
}

func main() {
	before := time.Now()

	numWorkers := 3
	numJobs := 10
	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create worker
	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	// Send values to tasks channel
	for i := range numJobs {
		tasks <- i
	}
	close(tasks)

	// Collect the results
	for range numJobs {
		result := <-results
		fmt.Println("Result:", result)
	}

	after := time.Now()
	fmt.Println("Time taken:", after.Sub(before))
}
