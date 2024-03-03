package kanban

import (
	"context"

	"github.com/google/uuid"
)

type Storer interface {
	Statuses(ctx context.Context) ([]Status, error)
	CardsByStatus(ctx context.Context, statusID uuid.UUID) ([]Card, error)
	Card(ctx context.Context, id uuid.UUID) (Card, error)
	UpsertCard(ctx context.Context, card Card) error
	DeleteCard(ctx context.Context, id uuid.UUID) error
}
