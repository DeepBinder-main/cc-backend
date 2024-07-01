-- name: DeleteFileStashURLByID :exec
DELETE FROM file_stash_url WHERE id = $1;
