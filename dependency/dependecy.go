package dependency

import (
	"github.com/Nihal1203/order-management-system/sqlc/namedqueries"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

func NewDependency(
	db *pgx.Conn,
	redis *redis.Client,
	queries *namedqueries.Queries,
) *AllDependency {

	return &AllDependency{
		DB:      db,
		Redis:   redis,
		Queries: queries,
	}
}

type AllDependency struct {
	DB      *pgx.Conn
	Redis   *redis.Client
	Queries *namedqueries.Queries
}
