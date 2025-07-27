package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigServer struct {
	Port        string
	API_VERSION string
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

	return &ConfigServer{
		Port:        port,
		API_VERSION: apiVersion,
	}, nil
}
