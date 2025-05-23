package main

import (
	"context"
	"fmt"
	"github.com/Lucas-Mol/go-studies/rssaggregator/internal/config"
	"github.com/Lucas-Mol/go-studies/rssaggregator/internal/database"
	"github.com/Lucas-Mol/go-studies/rssaggregator/internal/handlers"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	config.LoadingEnvFile()

	dbURL := config.GetEnvVar("DB_URL")
	conn := config.InitializeDBConn(dbURL)

	db := database.New(conn)
	apiCfg := handlers.ApiConfig{
		DB: db,
	}
	router := config.InitializeRouter(apiCfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go startScraping(ctx, db, 10, time.Minute)

	port := config.GetEnvVar("PORT")
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%v", port),
	}

	log.Printf("Server starting on port %v", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
