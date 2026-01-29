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

func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	//Get from post body
	filter := r.URL.Query().Get("fname")
	breweryFilter := r.URL.Query().Get("bid")

	var beers []BeerDTO
	var err error

	beers, err = getFilterResults(r.Context(), h.store, filter, breweryFilter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beers)
}

func getFilterResults(ctx context.Context, store BeerReadStore,
	filter string, breweryFilter string) ([]BeerDTO, error) {

	switch {
	case filter != "" && breweryFilter == "":
		return filterByName(ctx, store, filter)
	case breweryFilter != "" && filter == "":
		return filterByBrewery(ctx, store, breweryFilter)
	case breweryFilter == "" && filter == "":
		return store.List(ctx)
	default:
		return nil, errors.New("invalid filter combination. Choose either beer name (fname) or brewery id (bid).")
	}
}

func filterByName(ctx context.Context, store BeerReadStore,
	filter string) ([]BeerDTO, error) {

	if filter == "" {
		return nil, errors.New("filter cannot be empty")
	}

	beers, err := store.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	return beers, nil
}

func filterByBrewery(ctx context.Context, store BeerReadStore,
	breweryFilter string) ([]BeerDTO, error) {

	breweryId, err := strconv.ParseInt(breweryFilter, 10, 64)
	if err != nil {
		return nil, errors.New("invalid brewery id")
	}

	if breweryId <= 0 {
		return nil, errors.New("invalid brewery id")
	}

	beers, err := store.ListByBrewery(ctx, breweryId)
	if err != nil {
		return nil, err
	}

	return beers, nil
}
