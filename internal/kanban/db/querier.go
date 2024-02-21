// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	DeleteCard(ctx context.Context, id uuid.UUID) error
	GetCard(ctx context.Context, id uuid.UUID) (Card, error)
	ListCards(ctx context.Context) ([]Card, error)
	UpsertCard(ctx context.Context, arg UpsertCardParams) error
}

var _ Querier = (*Queries)(nil)