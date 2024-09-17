package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OptFunc func(*pgxpool.Config)

func WithMaxConn(maxConn int32) OptFunc {
	return func(cfg *pgxpool.Config) {
		cfg.MaxConns = maxConn
	}
}

func WithMinConn(minConn int32) OptFunc {
	return func(cfg *pgxpool.Config) {
		cfg.MinConns = minConn
	}
}

// NewConnection creates a new pooled connection for pgx
func NewConnection(
	ctx context.Context,
	url string,
	opts ...OptFunc,
) (
	*pgxpool.Pool,
	error,
) {
	// parse connection string as config struct
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string")
	}

	// pool settings defaults
	cfg.MaxConns = 10
	cfg.MinConns = 5
	cfg.MaxConnLifetime = time.Minute * 10
	cfg.MaxConnIdleTime = time.Minute * 5

	for _, fn := range opts {
		fn(cfg)
	}

	// connect
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping Postgres: %w", err)
	}

	return pool, nil
}

func BuildURL(url, dbName, user, pwd string) string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pwd, url, dbName)
}
