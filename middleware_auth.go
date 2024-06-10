package main

import (
	"fmt"
	"net/http"

	"github.com/luismingati/feedaggregator/internal/auth"
	"github.com/luismingati/feedaggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPiKey(r.Header)
		if err != nil {
			respondWitherror(w, http.StatusForbidden, fmt.Sprintf("error getting api key: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWitherror(w, http.StatusNotFound, fmt.Sprintf("could not find user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
