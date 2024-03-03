package kanban

import (
	"context"

	"github.com/google/uuid"
	"github.com/lokflam/simple-kanban/internal/kanban"
	"github.com/lokflam/simple-kanban/internal/store"
	"github.com/lokflam/simple-kanban/internal/store/kanban/postgresql"
)

type PostgreSQLStore struct {
	querier postgresql.Querier
}

var _ kanban.Storer = (*PostgreSQLStore)(nil)

func NewPostgreSQLStore(db postgresql.DBTX) *PostgreSQLStore {
	return &PostgreSQLStore{
		querier: postgresql.New(db),
	}
}

func (s *PostgreSQLStore) Statuses(ctx context.Context) ([]kanban.Status, error) {
	statuses, err := s.querier.Statuses(ctx)
	if err != nil {
		return nil, store.NewErrQueryFailed(err)
	}

	result := make([]kanban.Status, len(statuses))
	for i, status := range statuses {
		result[i] = kanban.Status{
			ID:   status.ID,
			Name: status.Name,
		}
	}

	return result, nil
}

func (s *PostgreSQLStore) CardsByStatus(ctx context.Context, statusID uuid.UUID) ([]kanban.Card, error) {
	cards, err := s.querier.CardsByStatus(ctx, statusID)
	if err != nil {
		return nil, store.NewErrQueryFailed(err)
	}

	result := make([]kanban.Card, len(cards))
	for i, card := range cards {
		result[i] = kanban.Card{
			ID:      card.Card.ID,
			Title:   card.Card.Title,
			Content: card.Card.Content,
			Status: kanban.Status{
				ID:   card.Status.ID,
				Name: card.Status.Name,
			},
		}
	}

	return result, nil
}

func (s *PostgreSQLStore) Card(ctx context.Context, id uuid.UUID) (kanban.Card, error) {
	card, err := s.querier.Card(ctx, id)
	if err != nil {
		return kanban.Card{}, store.NewErrQueryFailed(err)
	}

	return kanban.Card{
		ID:      card.Card.ID,
		Title:   card.Card.Title,
		Content: card.Card.Content,
		Status: kanban.Status{
			ID:   card.Status.ID,
			Name: card.Status.Name,
		},
	}, nil
}

func (s *PostgreSQLStore) UpsertCard(ctx context.Context, card kanban.Card) error {
	params := postgresql.UpsertCardParams{
		ID:       card.ID,
		Title:    card.Title,
		Content:  card.Content,
		StatusID: card.Status.ID,
	}

	if err := s.querier.UpsertCard(ctx, params); err != nil {
		return store.NewErrQueryFailed(err)
	}

	return nil
}

func (s *PostgreSQLStore) DeleteCard(ctx context.Context, id uuid.UUID) error {
	if err := s.querier.DeleteCard(ctx, id); err != nil {
		return store.NewErrQueryFailed(err)
	}

	return nil
}
