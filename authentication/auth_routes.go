package authentication

import (
	"github.com/gorilla/mux"
)

// api/v1/auth
func RegisterRoutes(router *mux.Router, authenticationManager *AuthenticationManager) {
	router.HandleFunc("/login", authenticationManager.loginHandler).Methods("POST")
}
