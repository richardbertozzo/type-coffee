package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewConnection creates a new pooled connection for pgx
func NewConnection(
	ctx context.Context,
	url string,
) (
	*pgxpool.Pool,
	error,
) {
	// parse connection string as config struct
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string")
	}

	// pool settings
	cfg.MaxConns = 10
	cfg.MinConns = 5
	cfg.MaxConnLifetime = time.Minute * 10
	cfg.MaxConnIdleTime = time.Minute * 5

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
