package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pranavkolte/chat-server-websocket/internal/config"
	postgresdb "github.com/pranavkolte/chat-server-websocket/internal/db/postgres/sqlc"
	"github.com/pranavkolte/chat-server-websocket/internal/handlers"
	"github.com/pranavkolte/chat-server-websocket/internal/managers"
	"github.com/pranavkolte/chat-server-websocket/internal/middleware"
	"github.com/pranavkolte/chat-server-websocket/internal/routes"
	"github.com/pranavkolte/chat-server-websocket/internal/util"
)

func main() {
	log.Println("Loading server configuration...")
	server_config, err := config.LoadConfigServer()
	if err != nil {
		log.Fatalf("Error loading server config: %v", err)
		return
	} else {
		log.Println("Server config loaded........")
	}

	util.SetJWTSecret(server_config.AUTHENTICATION.JWT_SECRET)

	log.Println("Setting up PostgreSQL DB")
	pgdb, err := sql.Open("postgres", server_config.POSTGRESQL_CONFIG.CONNECTION_URL)
	if err != nil {
		log.Fatalf("Failed to connect to DB \nerr: %v \n CONNECTION_URL: %v", err, server_config.POSTGRESQL_CONFIG.CONNECTION_URL)
	} else {
		log.Println("Postgres DB connection complete.....")
	}

	log.Println(server_config.POSTGRESQL_CONFIG.CONNECTION_URL)

	pgqueries := postgresdb.New(pgdb)

	// Create a new gorilla/mux router
	mainRouter := mux.NewRouter()

	// Create API subrouter
	apiRouter := mainRouter.PathPrefix("/api/" + server_config.API_VERSION).Subrouter()
	apiRouter.Use(middleware.JWTAuthMiddleware)

	// Authentication router
	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authenticationManager := managers.NewAuthenticationManager(pgqueries)
	authenticationHandler := &handlers.AuthenticationHandler{AuthenticationManager: authenticationManager}
	routes.AuthRoutes(authRouter, authenticationHandler)

	// User router
	userRouter := apiRouter.PathPrefix("/users").Subrouter()
	userManager := managers.NewUserManager(pgqueries)
	userHandler := &handlers.UserHandler{UserManager: userManager}
	routes.UserRoutes(userRouter, userHandler)

	// Start the server
	log.Printf("Server starting on port %v", server_config.SERVER_PORT)
	log.Fatal(http.ListenAndServe(":"+server_config.SERVER_PORT, mainRouter))
}
