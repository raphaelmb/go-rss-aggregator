package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/raphaelmb/go-rss/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondeWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondeWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not create feed: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondeWithError(w, http.StatusInternalServerError, fmt.Sprintf("Could not get feeds: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedsToFeeds(feeds))
}
