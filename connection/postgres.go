package connection

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type PostgresServer struct {
	ps *pgx.Conn
}

func NewPostgresServer() *PostgresServer {
	return &PostgresServer{}
}

func (p *PostgresServer) ConnectPostgresDB(context context.Context) (*pgx.Conn, error) {

	conn, err := pgx.Connect(context, "postgres://postgres:mypassword@localhost:5432/bank")
	if err != nil {
		return nil, err
	}
	log.Println("Postgres Connected Successfully")
	return conn, nil
}
