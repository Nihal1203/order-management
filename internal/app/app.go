package app

import (
	"context"
	"net/http"

	"github.com/Nihal1203/order-management-system/connection"
	"github.com/Nihal1203/order-management-system/dependency"
	"github.com/Nihal1203/order-management-system/sqlc/namedqueries"

	goahttp "goa.design/goa/v3/http"

	"github.com/daveamit/health"

	gorderservice "github.com/Nihal1203/order-management-system/gen/orders"
	guserservice "github.com/Nihal1203/order-management-system/gen/users"

	gorderhttp "github.com/Nihal1203/order-management-system/gen/http/orders/server"
	guserhttp "github.com/Nihal1203/order-management-system/gen/http/users/server"

	"github.com/Nihal1203/order-management-system/internal/middleware"
	implorder "github.com/Nihal1203/order-management-system/internal/orders"
	impluser "github.com/Nihal1203/order-management-system/internal/users"
)

func Run(ctx context.Context) error {

	health.EnsureService("postgres", "nihalns")

	pgInstance := connection.NewPostgresServer()
	pgconn, err := pgInstance.ConnectPostgresDB(ctx)
	if err != nil {
		return err
	}

	redisInstance := connection.NewRedisServer()
	redisconn, err := redisInstance.ConnectRedis(ctx)
	if err != nil {
		return err
	}

	queriesinstance := namedqueries.New(pgconn)

	depedency := dependency.NewDependency(pgconn, redisconn, queriesinstance)

	mux := goahttp.NewMuxer()
	handler := corsMiddleware(middleware.CookieMiddleware(mux))

	userService := impluser.NewUserService(depedency)
	userEndpoint := guserservice.NewEndpoints(userService)
	userHandler := guserhttp.New(userEndpoint, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
	guserhttp.Mount(mux, userHandler)

	orderSerivce := implorder.NewOrderService(depedency)
	orderEndpoint := gorderservice.NewEndpoints(orderSerivce)
	ordeHandler := gorderhttp.New(orderEndpoint, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)
	gorderhttp.Mount(mux, ordeHandler)
	return http.ListenAndServe(":3000", handler)

}

func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
