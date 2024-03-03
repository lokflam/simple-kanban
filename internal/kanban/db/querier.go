// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	Card(ctx context.Context, id uuid.UUID) (Card, error)
	CardsByStatus(ctx context.Context, statusID uuid.UUID) ([]Card, error)
	DeleteCard(ctx context.Context, id uuid.UUID) error
	Statuses(ctx context.Context) ([]StatusesRow, error)
	UpsertCard(ctx context.Context, arg UpsertCardParams) error
}

var _ Querier = (*Queries)(nil)
