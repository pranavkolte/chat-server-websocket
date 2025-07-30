package managers

import (
	"context"

	postgresdb "github.com/pranavkolte/chat-server-websocket/internal/db/postgres/sqlc"
)

type UserManager struct {
	PgQueries *postgresdb.Queries
}

func NewUserManager(pgQueries *postgresdb.Queries) *UserManager {
	return &UserManager{
		PgQueries: pgQueries,
	}
}

func (userManager *UserManager) GetUsersPaginated(ctx context.Context, args postgresdb.GetUsersPaginatedParams) ([]postgresdb.GetUsersPaginatedRow, error) {
	return userManager.PgQueries.GetUsersPaginated(ctx, args)
}
