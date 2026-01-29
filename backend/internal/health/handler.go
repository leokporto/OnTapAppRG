package health

import (
	"net/http"
)

type Handler struct {
	healthMessage string
}

func NewHandler() *Handler {
	return &Handler{healthMessage: "Service is healthy!"}
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.healthMessage))
}
