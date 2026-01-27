package beerstyle

import (
	"context"
	"encoding/json"
	"net/http"
)

type Handler struct {
	store Store
}

func NewHandler(store Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	styles, err := listBeerStyles(r.Context(), h.store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(styles)
}

func listBeerStyles(
	ctx context.Context,
	store Store,
) ([]BeerStyle, error) {

	styles, err := store.List(ctx)
	if err != nil {
		return nil, err
	}

	return styles, nil
}
