package managers

import (
	"context"

	postgresdb "github.com/pranavkolte/chat-server-websocket/internal/db/postgres/sqlc"
)

type AuthenticationManager struct {
	PgQueries *postgresdb.Queries
}

func NewAuthenticationManager(pgQueries *postgresdb.Queries) *AuthenticationManager {
	return &AuthenticationManager{
		PgQueries: pgQueries,
	}
}

func (authManager *AuthenticationManager) CreateUser(ctx context.Context, params postgresdb.CreateUserParams) (postgresdb.User, error) {
	return authManager.PgQueries.CreateUser(ctx, params)
}

func (authManager *AuthenticationManager) GetUserByEmail(ctx context.Context, email string) (postgresdb.User, error) {
	return authManager.PgQueries.GetUserByEmail(ctx, email)
}
