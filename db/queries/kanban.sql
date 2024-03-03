-- name: Statuses :many
SELECT status.*
FROM status
LEFT JOIN board_status ON status.id = board_status.status_id
ORDER BY board_status.position;

-- name: CardsByStatus :many
SELECT sqlc.embed(card), sqlc.embed(status)
FROM card
JOIN status ON card.status_id = status.id
WHERE status.id = $1
ORDER BY card.created_at DESC;

-- name: Card :one
SELECT sqlc.embed(card), sqlc.embed(status)
FROM card
JOIN status ON card.status_id = status.id
WHERE card.id = $1;

-- name: UpsertCard :exec
INSERT INTO
    card (id, title, content, status_id, created_at, updated_at)
VALUES
    ($1, $2, $3, $4, $5, $6)
ON CONFLICT (id) DO UPDATE
SET
    title = $2,
    content = $3,
    status_id = $4,
    updated_at = $6;

-- name: DeleteCard :exec
DELETE FROM card
WHERE id = $1;
