package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/raphaelmb/go-rss/internal/auth"
	"github.com/raphaelmb/go-rss/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondeWithError(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				respondeWithError(w, http.StatusNotFound, fmt.Sprintf("No user found: %v", err))
				return

			}
			respondeWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
