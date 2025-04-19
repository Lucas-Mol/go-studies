package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Lucas-Mol/go-studies/rssaggregator/auth"
	"github.com/Lucas-Mol/go-studies/rssaggregator/internal/database"
	"github.com/Lucas-Mol/go-studies/rssaggregator/models"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (apiCfg ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
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

	newUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
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

func (apiCfg ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusForbidden, err.Error())
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusForbidden, fmt.Sprintf("Couldn't get user: %v", err.Error()))
		return
	}

	respondWithJSON(w, http.StatusOK, models.DatabaseUserToUser(user))
}
