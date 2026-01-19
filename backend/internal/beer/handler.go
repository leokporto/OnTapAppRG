package beer

import (
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
	beerId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || beerId <= 0 {
		errMsg := errors.New("Id is invalid")
		w.Write([]byte(errMsg.Error()))
		return
	}

	response := Beer{ID: 1, Name: "Punk IPA", StyleID: 1, BreweryID: 1, FullName: "Punk IPA", ABV: 6.5, MinIBU: 45, MaxIBU: 55}

	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetStyles(w http.ResponseWriter, r *http.Request) {
	styles := []string{"IPA", "Lager"}

	json.NewEncoder(w).Encode(styles)
}
