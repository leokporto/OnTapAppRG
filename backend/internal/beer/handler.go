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

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	beers, err := h.store.ListAll(r.Context())
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
	store BeerStore,
	id int64,
) (BeerResponse, error) {

	if id <= 0 {
		return BeerResponse{}, errors.New("invalid id")
	}

	beer, err := store.GetByID(ctx, id)
	if err != nil {
		return BeerResponse{}, err
	}

	return beer, nil
}

func (h *Handler) GetStyles(w http.ResponseWriter, r *http.Request) {
	styles, err := getBeerStyles(r.Context(), h.store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(styles)
}

func getBeerStyles(
	ctx context.Context,
	store BeerStore,
) ([]BeerStyle, error) {

	styles, err := store.GetStyles(ctx)
	if err != nil {
		return nil, err
	}

	return styles, nil
}

func (h *Handler) GetBreweries(w http.ResponseWriter, r *http.Request) {
	breweries, err := getBreweries(r.Context(), h.store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(breweries)
}

func getBreweries(
	ctx context.Context,
	store BeerStore,
) ([]Brewery, error) {

	breweries, err := store.GetBreweries(ctx)
	if err != nil {
		return nil, err
	}

	return breweries, nil
}
