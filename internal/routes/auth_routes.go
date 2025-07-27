package routes

import (
	"github.com/gorilla/mux"
	"github.com/pranavkolte/chat-server-websocket/internal/handlers"
)

// api/v1/auth
func RegisterRoutes(router *mux.Router, authenticationHandler *handlers.AuthenticationHandler) {
	router.HandleFunc("/login", authenticationHandler.LoginHandler).Methods("POST")
}
