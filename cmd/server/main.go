package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pranavkolte/chat-server-websocket/internal/config"
	"github.com/pranavkolte/chat-server-websocket/internal/handlers"
	"github.com/pranavkolte/chat-server-websocket/internal/managers"
	"github.com/pranavkolte/chat-server-websocket/internal/routes"
)

func main() {
	log.Println("Loading server configuration...")
	server_config, err := config.LoadConfigServer()
	if err != nil {
		log.Fatalf("Error loading server config: %v", err)
		return
	}

	// Create a new gorilla/mux router
	mainRouter := mux.NewRouter()

	// Create API subrouter
	apiRouter := mainRouter.PathPrefix("/api/" + server_config.API_VERSION).Subrouter()

	// Authentication router
	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authenticationManager := managers.NewAuthenticationManager()
	authenticationHandler := &handlers.AuthenticationHandler{AuthenticationManager: authenticationManager}
	routes.RegisterRoutes(authRouter, authenticationHandler)

	// Start the server
	log.Printf("Server starting on port %v", server_config.Port)
	log.Fatal(http.ListenAndServe(":"+server_config.Port, mainRouter))
}
