package authentication

import (
	"fmt"
	"net/http"
)

func (authenticationManager *AuthenticationManager) loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Login endpoint reached successfully"}`)
}
