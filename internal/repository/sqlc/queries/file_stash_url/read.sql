-- name: GetFileStashURLByID :one
SELECT id, url, created_at FROM file_stash_url WHERE id = $1;

-- name: ListFileStashURLs :many
SELECT id, url, created_at FROM file_stash_url ORDER BY created_at DESC;
