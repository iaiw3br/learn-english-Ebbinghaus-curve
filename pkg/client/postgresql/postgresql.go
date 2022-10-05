package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"tg-bot-learning-english/internal/config"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, p config.Postgresql) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("%s://%s:%s@%s:%s/%s", p.Driver, p.Username, p.Password, p.Host, p.Port, p.Database)

	pool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
