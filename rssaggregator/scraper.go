package main

import (
	"context"
	"database/sql"
	"github.com/Lucas-Mol/go-studies/rssaggregator/internal/database"
	"github.com/google/uuid"
	"log"
	"strings"
	"sync"
	"time"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("erro fetching feeds: ", err.Error())
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scraperFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scraperFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	log.Printf("Scraping feed %v", feed.ID)

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("erro marking feed as fetched: ", err.Error())
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("error fetching feed from url: ", err.Error())
		return
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

	for _, item := range rssFeed.Channel.Item {
		itemDescription := sql.NullString{}
		if item.Description != "" {
			itemDescription.String = item.Description
			itemDescription.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Couldn't parse date %v with error: %v", item.PubDate, err.Error())
			continue
		}

		post, err := db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Description: itemDescription,
				PublishedAt: pubAt,
				Url:         item.Link,
				FeedID:      feed.ID,
			})
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key") {
				log.Println("Failed to create post: ", err.Error())
			}
			continue
		}

		log.Printf("Post %v with %v created for feed %v", post.ID, post.Title, feed.Name)
	}
}
