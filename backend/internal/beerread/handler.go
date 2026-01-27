package beerread

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store BeerReadStore
}

func NewHandler(store BeerReadStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	beers, err := h.store.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beers)

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
	store BeerReadStore,
	id int64,
) (BeerDTO, error) {
	if id <= 0 {
		return BeerDTO{}, errors.New("invalid id")
	}

	beer, err := store.GetById(ctx, id)
	if err != nil {
		return BeerDTO{}, err
	}

	return beer, nil
}
