package beer

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	beers := []Beer{
		{ID: 1, Name: "Punk IPA", Style: "IPA", Brewery: "Punk", FullName: "Punk IPA", ABV: 6.5, MinIBU: 45, MaxIBU: 55},
		{ID: 2, Name: "Heineken", Style: "Lager", Brewery: "Heineken", FullName: "Heineken Lager", ABV: 5, MinIBU: 10, MaxIBU: 11},
	}

	json.NewEncoder(w).Encode(beers)

}

func GetById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	beerId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || beerId <= 0 {
		errMsg := errors.New("Id is invalid")
		w.Write([]byte(errMsg.Error()))
		return
	}

	response := Beer{ID: 1, Name: "Punk IPA", Style: "IPA", Brewery: "Punk", FullName: "Punk IPA", ABV: 6.5, MinIBU: 45, MaxIBU: 55}

	json.NewEncoder(w).Encode(response)
}

func GetStyles(w http.ResponseWriter, r *http.Request) {
	styles := []string{"IPA", "Lager"}

	json.NewEncoder(w).Encode(styles)
}
