-- name: ListCards :many
SELECT
    *
FROM
    cards;

-- name: GetCard :one
SELECT
    *
FROM
    cards
WHERE
    id = $1;

-- name: UpsertCard :exec
INSERT INTO
    cards (id, title, content, status_id, created_at, updated_at)
VALUES
    ($1, $2, $3, '018dbc48-4899-7aac-a1fa-0680a50c82a9', $4, $5)
ON CONFLICT (id)
DO UPDATE
    SET
        title = $2,
        content = $3,
        updated_at = $5;

-- name: DeleteCard :exec
DELETE FROM
    cards
WHERE
    id = $1;
