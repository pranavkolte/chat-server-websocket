package routes

import (
	"github.com/gorilla/mux"
	"github.com/pranavkolte/chat-server-websocket/internal/handlers"
)

// api/v1/users
func UserRoutes(router *mux.Router, userHandler *handlers.UserHandler) {
	router.HandleFunc("", userHandler.GetUsersPaginatedHandler).Methods("GET")
}
