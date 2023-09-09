-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id, last_fetched_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextsFeedToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedFeteched :one
UPDATE feeds
SET updated_at = CURRENT_TIMESTAMP AT TIME ZONE 'UTC', last_fetched_at = CURRENT_TIMESTAMP AT TIME ZONE 'UTC'
WHERE id = $1
RETURNING *;
