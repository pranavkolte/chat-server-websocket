package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	postgresdb "github.com/pranavkolte/chat-server-websocket/internal/db/postgres/sqlc"
	"github.com/pranavkolte/chat-server-websocket/internal/managers"
	"github.com/pranavkolte/chat-server-websocket/internal/schemas"
)

type AuthenticationHandler struct {
	AuthenticationManager *managers.AuthenticationManager
}

func (authHandler *AuthenticationHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Login endpoint reached successfully"}`)
}

func (authHanlder *AuthenticationHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req schemas.SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID := uuid.New()

	params := postgresdb.CreateUserParams{
		UserID:   userID,
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	pgUser, err := authHanlder.AuthenticationManager.CreateUser(r.Context(), params)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"user_id": "%v", "username": "%v"}`, pgUser.UserID, pgUser.Username)
}
