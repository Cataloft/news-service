package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"log"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

type Postgres struct {
	Conn *pgxpool.Pool
	DB   *reform.DB
}

func New(dbURL string) *Postgres {
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
	}
	sqlDBPool := stdlib.OpenDBFromPool(pool)

	db := reform.NewDB(sqlDBPool, postgresql.Dialect, reform.NewPrintfLogger(log.Printf))

	return &Postgres{
		Conn: pool,
		DB:   db,
	}
}
