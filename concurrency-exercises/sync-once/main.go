package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Just once called")
}

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutine #", i)
			once.Do(initialize)
		}()
	}

	wg.Wait()
}
