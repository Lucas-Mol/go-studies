package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numGoRoutines := 10
	wg.Add(numGoRoutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numGoRoutines {
		go increment()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}
