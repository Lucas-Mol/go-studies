package main

import (
	"context"
	"fmt"
	"time"
)

const MaxConcurrency = 5

var urls = []string{
	"https://golang.org/",
	"https://pkg.go.dev/std",
	"https://golang.org/pkg/fmt",
	"https://golang.org/pkg/os/",
	"https://golang.org/LICENSE",
	"https://httpbin.org/delay/2",
	"https://httpbin.org/status/200",
	"https://httpbin.org/delay/1",
	"https://httpbin.org/delay/2",
	"https://httpbin.org/status/200",
	"https://httpbin.org/delay/1",
	"https://golang.org/",
	"https://pkg.go.dev/std",
	"https://golang.org/pkg/fmt",
	"https://golang.org/pkg/os/",
	"https://golang.org/LICENSE",
	"https://httpbin.org/delay/2",
	"https://httpbin.org/status/200",
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	before := time.Now()
	err := crawlURLs(ctx, urls, MaxConcurrency)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Crawler finished successfully.")
	}
	after := time.Now()
	duration := after.Sub(before)
	fmt.Printf("Duration: %v\n", duration)
}
