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

func (cfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := &parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	newUser, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %v", err))
	}

	respondWithJSON(w, http.StatusCreated, models.DatabaseUserToUser(newUser))
}

func (cfg *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, models.DatabaseUserToUser(user))
}

func (cfg *ApiConfig) HandlerGetPostsByUserFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := cfg.DB.GetPostsForUser(
		r.Context(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  10,
		},
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,
			fmt.Sprintf("Couldn't get posts: %v", err.Error()))
		return
	}

	respondWithJSON(w, http.StatusOK, models.DatabasePostsToPosts(posts))

}
