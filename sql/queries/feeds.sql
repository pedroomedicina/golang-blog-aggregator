-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name, feeds.url, u.name as user_name
FROM feeds
INNER JOIN users u ON feeds.user_id = u.id;

-- name: GetFeedByUrl :one
SELECT feeds.id, feeds.name, feeds.url, u.name as user_name
FROM feeds
INNER JOIN users u ON feeds.user_id = u.id
WHERE feeds.url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at NULLS FIRST, updated_at
LIMIT 1;