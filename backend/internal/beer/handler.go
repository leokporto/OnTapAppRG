package beer

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store BeerStore
}

func NewHandler(store BeerStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Id is invalid", http.StatusBadRequest)
		return
	}

	beer, err := getBeerByID(r.Context(), h.store, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beer)
}

func getBeerByID(
	ctx context.Context,
	store BeerStore,
	id int64,
) (Beer, error) {

	if id <= 0 {
		return Beer{}, errors.New("invalid id")
	}

	beer, err := store.GetById(ctx, id)
	if err != nil {
		return Beer{}, err
	}

	return beer, nil
}
