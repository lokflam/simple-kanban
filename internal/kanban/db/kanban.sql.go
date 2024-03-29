// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: kanban.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const card = `-- name: Card :one
SELECT id, title, content, status_id, created_at, updated_at
FROM card
WHERE id = $1
`

func (q *Queries) Card(ctx context.Context, id uuid.UUID) (Card, error) {
	row := q.db.QueryRow(ctx, card, id)
	var i Card
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.StatusID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const cardsByStatus = `-- name: CardsByStatus :many
SELECT id, title, content, status_id, created_at, updated_at
FROM card
WHERE status_id = $1
ORDER BY created_at DESC
`

func (q *Queries) CardsByStatus(ctx context.Context, statusID uuid.UUID) ([]Card, error) {
	rows, err := q.db.Query(ctx, cardsByStatus, statusID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Card
	for rows.Next() {
		var i Card
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.StatusID,
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

const deleteCard = `-- name: DeleteCard :exec
DELETE FROM card
WHERE id = $1
`

func (q *Queries) DeleteCard(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteCard, id)
	return err
}

const statuses = `-- name: Statuses :many
SELECT status.id, status.name, status.created_at, status.updated_at, board_status.id, board_status.status_id, board_status.position, board_status.created_at, board_status.updated_at
FROM status
LEFT JOIN board_status ON status.id = board_status.status_id
ORDER BY board_status.position
`

type StatusesRow struct {
	Status      Status
	BoardStatus BoardStatus
}

func (q *Queries) Statuses(ctx context.Context) ([]StatusesRow, error) {
	rows, err := q.db.Query(ctx, statuses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StatusesRow
	for rows.Next() {
		var i StatusesRow
		if err := rows.Scan(
			&i.Status.ID,
			&i.Status.Name,
			&i.Status.CreatedAt,
			&i.Status.UpdatedAt,
			&i.BoardStatus.ID,
			&i.BoardStatus.StatusID,
			&i.BoardStatus.Position,
			&i.BoardStatus.CreatedAt,
			&i.BoardStatus.UpdatedAt,
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
