package main

import (
	"fmt"
	"sync"
	"time"
)

// ======= BASIC EXAMPLE USING CHANNELS =======

// func worker(id int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("WorkerID %d starting\n", id)
// 	time.Sleep(time.Second) //simulate some work
// 	results <- id * 2
// 	fmt.Printf("WorkerID %d finished\n", id)
// }

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d starting\n", id)
	time.Sleep(time.Second) //simulate some work
	for task := range tasks {
		results <- task * 2
	}
	fmt.Printf("WorkerID %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3
	numJobs := 5
	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker(i, tasks, results, &wg)
	}

	for i := range numJobs {
		tasks <- i + 1
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}
}

// =========== BASIC EXAMPLE WITHOUT USING CHANNELS =============
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// wg.Add(1) WRONG PRACTICE!!!
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second) // simulate some time spent on processing the task
// 	fmt.Printf("Worker %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers) // THIS IS THE CORRECT WAY OF ADDING COUNTER TO WAIT GROUP

// 	// Launch workers
// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait() // blocking mechanism
// 	fmt.Println("All workers finished")
// }
