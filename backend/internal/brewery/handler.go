package brewery

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store Store
}

func NewHandler(store Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	breweries, err := listBreweries(r.Context(), h.store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(breweries)
}

func listBreweries(
	ctx context.Context,
	store Store,
) ([]Brewery, error) {

	breweries, err := store.List(ctx)
	if err != nil {
		return nil, err
	}

	return breweries, nil
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Id is invalid", http.StatusBadRequest)
		return
	}

	brewery, err := getBreweryByID(r.Context(), h.store, id)
	if err != nil {
		http.Error(w, "Err on handler: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brewery)
}

func getBreweryByID(
	ctx context.Context,
	store Store,
	id int64,
) (Brewery, error) {

	if id <= 0 {
		return Brewery{}, errors.New("invalid id")
	}

	brewery, err := store.GetById(ctx, id)
	if err != nil {
		return Brewery{}, err
	}

	return brewery, nil
}
