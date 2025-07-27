package handlers

import (
	"fmt"
	"net/http"

	"github.com/pranavkolte/chat-server-websocket/internal/managers"
)

type AuthenticationHandler struct {
	AuthenticationManager *managers.AuthenticationManager
}

func (authenticationHandler *AuthenticationHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Login endpoint reached successfully"}`)
}
