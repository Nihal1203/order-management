package connection

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisServer struct {
	rs *redis.Client
}

func NewRedisServer() *RedisServer {
	return &RedisServer{}
}

func (r *RedisServer) ConnectRedis(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		// Username: "default",
		// Password: "nihal-db",
		// DB:       0,
		// Protocol: 2,
		Addr: "localhost:6379",
		DB:   0,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	log.Println("Redis Connected Successfully")

	return rdb, nil
}
