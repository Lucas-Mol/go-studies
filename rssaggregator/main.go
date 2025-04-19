package main

import (
	"fmt"
	"github.com/Lucas-Mol/go-studies/rssaggregator/config"
	"github.com/Lucas-Mol/go-studies/rssaggregator/handlers"
	"github.com/Lucas-Mol/go-studies/rssaggregator/internal/database"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	config.LoadingFromEnvFile()

	dbURL := config.GetEnvVar("DB_URL")
	conn := config.InitializeDBConn(dbURL)

	apiCfg := handlers.ApiConfig{
		DB: database.New(conn),
	}

	port := config.GetEnvVar("PORT")
	router := config.InitializeRouter(apiCfg)

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
