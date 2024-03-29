// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: kanban.sql

package postgresql

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const card = `-- name: Card :one
SELECT card.id, card.title, card.content, card.status_id, card.created_at, card.updated_at, status.id, status.name, status.created_at, status.updated_at
FROM card
JOIN status ON card.status_id = status.id
WHERE card.id = $1
`

type CardRow struct {
	Card   Card
	Status Status
}

func (q *Queries) Card(ctx context.Context, id uuid.UUID) (CardRow, error) {
	row := q.db.QueryRow(ctx, card, id)
	var i CardRow
	err := row.Scan(
		&i.Card.ID,
		&i.Card.Title,
		&i.Card.Content,
		&i.Card.StatusID,
		&i.Card.CreatedAt,
		&i.Card.UpdatedAt,
		&i.Status.ID,
		&i.Status.Name,
		&i.Status.CreatedAt,
		&i.Status.UpdatedAt,
	)
	return i, err
}

const cardsByStatus = `-- name: CardsByStatus :many
SELECT card.id, card.title, card.content, card.status_id, card.created_at, card.updated_at, status.id, status.name, status.created_at, status.updated_at
FROM card
JOIN status ON card.status_id = status.id
WHERE status.id = $1
ORDER BY card.created_at DESC
`

type CardsByStatusRow struct {
	Card   Card
	Status Status
}

func (q *Queries) CardsByStatus(ctx context.Context, id uuid.UUID) ([]CardsByStatusRow, error) {
	rows, err := q.db.Query(ctx, cardsByStatus, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CardsByStatusRow
	for rows.Next() {
		var i CardsByStatusRow
		if err := rows.Scan(
			&i.Card.ID,
			&i.Card.Title,
			&i.Card.Content,
			&i.Card.StatusID,
			&i.Card.CreatedAt,
			&i.Card.UpdatedAt,
			&i.Status.ID,
			&i.Status.Name,
			&i.Status.CreatedAt,
			&i.Status.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteCard = `-- name: DeleteCard :exec
DELETE FROM card
WHERE id = $1
`

func (q *Queries) DeleteCard(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteCard, id)
	return err
}

const statuses = `-- name: Statuses :many
SELECT status.id, status.name, status.created_at, status.updated_at
FROM status
LEFT JOIN board_status ON status.id = board_status.status_id
ORDER BY board_status.position
`

func (q *Queries) Statuses(ctx context.Context) ([]Status, error) {
	rows, err := q.db.Query(ctx, statuses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Status
	for rows.Next() {
		var i Status
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const upsertCard = `-- name: UpsertCard :exec
INSERT INTO
    card (id, title, content, status_id, created_at, updated_at)
VALUES
    ($1, $2, $3, $4, $5, $6)
ON CONFLICT (id) DO UPDATE
SET
    title = $2,
    content = $3,
    status_id = $4,
    updated_at = $6
`

type UpsertCardParams struct {
	ID        uuid.UUID
	Title     string
	Content   string
	StatusID  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) UpsertCard(ctx context.Context, arg UpsertCardParams) error {
	_, err := q.db.Exec(ctx, upsertCard,
		arg.ID,
		arg.Title,
		arg.Content,
		arg.StatusID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}
