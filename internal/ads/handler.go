package ads

import (
	"log"
	"net/http"

	"github.com/NavroO/adhub/internal/shared"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.List)
	return r
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	entries, err := h.svc.List(r.Context())
	if err != nil {
		log.Printf("❌ Failed to load ads: %v", err)
		shared.RespondError(w, http.StatusInternalServerError, "Failed to load ads")
		return
	}
	shared.RespondJSON(w, http.StatusOK, entries)
}
