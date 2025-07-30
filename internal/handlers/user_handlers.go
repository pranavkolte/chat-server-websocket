package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	postgresdb "github.com/pranavkolte/chat-server-websocket/internal/db/postgres/sqlc"
	"github.com/pranavkolte/chat-server-websocket/internal/managers"
)

type UserHandler struct {
	UserManager *managers.UserManager
}

func (userHandler *UserHandler) GetUsersPaginatedHandler(w http.ResponseWriter, r *http.Request) {

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10 // default limit
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	paginationParams := postgresdb.GetUsersPaginatedParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	users, err := userHandler.UserManager.GetUsersPaginated(r.Context(), paginationParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
