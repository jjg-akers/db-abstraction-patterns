package framework

import (
	"context"
	"database/sql"
)

// Beginner... begines transactions
type beginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	contextExecuter
}

// ContextExecuter... can perform SQL queries with context
type contextExecuter interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}
