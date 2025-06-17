package main

import (
	"fmt"
	"time"
)

const NumWorkers = 3

type TicketRequest struct {
	personID  int
	numTicket int
	cost      int
}

// Simulate processing fo ticket requests
func ticketProcessor(requests <-chan TicketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf(
			"Processing %d ticket(s) of personID %d with total cost %d\n",
			req.personID,
			req.personID,
			req.cost,
		)
		// Simulate processing time
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func main() {
	before := time.Now()
	numRequests := 5
	price := 15
	ticketRequests := make(chan TicketRequest, numRequests)
	ticketResults := make(chan int, numRequests)

	// Start ticket workers
	for range NumWorkers {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// Send ticket requests
	for i := range numRequests {
		id := i + 1
		ticketRequests <- TicketRequest{
			id,
			id * 2,
			id * price,
		}
	}
	close(ticketRequests)

	// Collect results
	for range numRequests {
		fmt.Printf("Ticket for personID %d processed successfully!\n", <-ticketResults)
	}

	after := time.Now()
	fmt.Printf("Time taken: %v\n", after.Sub(before))
}
