package shared

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	log.Error().Int("status", code).Msg("❌ " + message)

	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(ErrorResponse{Error: message})
	if err != nil {
		log.Error().Err(err).Msg("❌ failed to encode error response")
	}
}
