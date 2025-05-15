package main

import (
	"context"
	"fmt"
	"net/http"
)

func crawlURLs(ctx context.Context, urls []string, maxConcurrent int) error {
	type result struct {
		url string
		err error
	}

	sem := make(chan struct{}, maxConcurrent)
	results := make(chan result, len(urls))

	for _, url := range urls {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sem <- struct{}{}:
			go func(u string) {
				defer func() { <-sem }() //returns a buffered channel slot as available

				err := visitURL(ctx, u)
				results <- result{url: u, err: err}
			}(url)
		}
	}

	for range len(urls) {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case res := <-results:
			if res.err != nil {
				fmt.Printf("Error in %s: %v\n", res.url, res.err)
				return fmt.Errorf("one or more URLs failed")
			} else {
				fmt.Printf("Success: %s\n", res.url)
			}
		}
	}

	return nil
}

func visitURL(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	return nil
}
