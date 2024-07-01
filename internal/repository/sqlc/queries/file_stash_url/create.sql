-- name: CreateFileStashURL :exec
INSERT INTO file_stash_url (url) VALUES ($1);
