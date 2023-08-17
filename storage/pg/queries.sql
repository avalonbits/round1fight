-- name: GetPerson :one
SELECT id, nickname, name, birthday, stack FROM Person where id = $1;

-- name: CountPerson :one
SELECT COUNT(*) person_count from Person;

-- name: CreatePerson :exec
INSERT INTO Person (id, nickname, name, birthday, stack) VALUES ($1, $2, $3, $4, $5);

-- name: SearchPerson :many
SELECT id, nickname, name, birthday, stack
    FROM Person WHERE docsearch @@ plainto_tsquery(@query::text) LIMIT 50;
