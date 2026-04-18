package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/username/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerPolkaWebhooks(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil || apiKey != cfg.PolkaKey {
		respondWithError(w, http.StatusUnauthorized, "Invalid API key")
		return
	}

	// Dynamically map the nested JSON structure Polka mathematically sends us!
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// 1. Is it a generic irrelevant notification? Throw it securely in the trash with a 204.
	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// 2. We dynamically snag the UserID payload and execute the DB upgrade locally
	userID, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID format")
		return
	}

	err = cfg.DB.UpgradeUserToChirpyRed(r.Context(), userID)
	if err != nil {
		// If SQL throws a rejection here natively, it physically couldn't map the user inside the Database natively
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// 3. We successfully granted premium access mechanically, shoot Polka a 2XX Success Signal!
	w.WriteHeader(http.StatusNoContent)
}
