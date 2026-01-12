package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leokporto/OnTapAppRG/backend/internal/health"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", logging(health.Handler()))

	http.ListenAndServe(":8080", r)
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}
