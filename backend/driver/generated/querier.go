// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db_driver

import (
	"context"
)

type Querier interface {
	FindTodo(ctx context.Context) ([]TaskCanvasTodo, error)
}

var _ Querier = (*Queries)(nil)