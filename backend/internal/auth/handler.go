package auth

import (
	"encoding/json"
	"net/http"

	"google.golang.org/api/idtoken"
)

type Handler struct {
	google  *GoogleConfig
	service *AuthService
	jwt     *JWTService
}

func NewHandler(
	google *GoogleConfig,
	service *AuthService,
	jwt *JWTService,
) *Handler {
	return &Handler{
		google:  google,
		service: service,
		jwt:     jwt,
	}
}

func (h *Handler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := h.google.OAuth.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusFound)
}

func (h *Handler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "missing code", http.StatusBadRequest)
		return
	}

	token, err := h.google.OAuth.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "token exchange failed", http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "missing id_token", http.StatusUnauthorized)
		return
	}

	payload, err := idtoken.Validate(ctx, rawIDToken, h.google.OAuth.ClientID)
	if err != nil {
		http.Error(w, "invalid id_token", http.StatusUnauthorized)
		return
	}

	subject, _ := payload.Claims["sub"].(string)
	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)

	u, err := h.service.LoginWithOpenID(
		ctx,
		"google",
		subject,
		email,
		name,
	)
	if err != nil {
		http.Error(w, "login failed", http.StatusInternalServerError)
		return
	}

	jwttoken, err := h.jwt.Generate(
		u.ID,
		u.Email,
		"user",
	)
	if err != nil {
		http.Error(w, "token generation failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": jwttoken,
	})

}
