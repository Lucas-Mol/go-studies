package main

import (
	"fmt"
	"time"
)

type StatefulWorker struct {
	count int
	ch    chan int
}

func (sw *StatefulWorker) Start() {
	go func() {
		for {
			select {
			case value := <-sw.ch:
				sw.count += value
				fmt.Println("Current count:", sw.count)
			}
		}
	}()
}

func (sw *StatefulWorker) Send(value int) {
	sw.ch <- value
}

func main() {
	statefulWorker := &StatefulWorker{ch: make(chan int)}

	statefulWorker.Start()

	for i := range 5 {
		statefulWorker.Send(i)
		time.Sleep(500 * time.Millisecond)
	}
}
