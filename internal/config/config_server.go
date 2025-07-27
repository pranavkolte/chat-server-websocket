package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigServer struct {
	SERVER_PORT       string
	API_VERSION       string
	DATA_SOURCE_NAME  string
	POSTGRESQL_CONFIG ConfigPostgreSQL
}

type ConfigPostgreSQL struct {
	HOST           string
	PORT           string
	DB_NAME        string
	USER           string
	PASSWORD       string
	CONNECTION_URL string
}

func LoadConfigServer() (*ConfigServer, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return nil, err
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Println("SERVER_PORT not set, using default port 8080")
		port = "8080"
	}

	apiVersion := os.Getenv("API_VERSION")
	if apiVersion == "" {
		log.Println("API_VERSION not set, using default version v1")
		apiVersion = "v1"
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresDBName := os.Getenv("POSTGRES_DB_NAME")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	sslMode := os.Getenv("SSL_MODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	connectionURL := "postgres://" + postgresUser + ":" + postgresPassword + "@" + postgresHost + ":" + postgresPort + "/" + postgresDBName + "?sslmode=" + sslMode

	postgresConfig := ConfigPostgreSQL{
		HOST:           postgresHost,
		PORT:           postgresPort,
		DB_NAME:        postgresDBName,
		USER:           postgresUser,
		PASSWORD:       postgresPassword,
		CONNECTION_URL: connectionURL,
	}
	return &ConfigServer{
		SERVER_PORT:       port,
		API_VERSION:       apiVersion,
		POSTGRESQL_CONFIG: postgresConfig,
	}, nil
}
