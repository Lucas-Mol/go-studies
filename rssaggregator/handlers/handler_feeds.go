package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Lucas-Mol/go-studies/rssaggregator/internal/database"
	"github.com/Lucas-Mol/go-studies/rssaggregator/models"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := &parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create feed: %v", err))
	}

	respondWithJSON(w, http.StatusCreated, models.DatabaseFeedToFeed(feed))
}

func (cfg *ApiConfig) HandlerGetAllFeedsByUser(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get all feeds: %v", err))
	}

	respondWithJSON(w, http.StatusCreated, models.DatabaseFeedsToFeeds(feeds))
}
