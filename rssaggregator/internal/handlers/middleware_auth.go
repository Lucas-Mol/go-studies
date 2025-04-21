package handlers

import (
	"fmt"
	"github.com/Lucas-Mol/go-studies/rssaggregator/internal/auth"
	"github.com/Lucas-Mol/go-studies/rssaggregator/internal/database"
	"net/http"
)

type AuthedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) MiddlewareAuth(handler AuthedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, err.Error())
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusForbidden, fmt.Sprintf("Couldn't get user: %v", err.Error()))
			return
		}

		handler(w, r, user)
	}
}
