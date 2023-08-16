-- name: GetPerson :one
SELECT * FROM Person where id = $1;

-- name: CountPerson :one
SELECT COUNT(*) person_count from Person;

-- name: CreatePerson :exec
INSERT INTO Person (id, nickname, name, birthday, stack) VALUES ($1, $2, $3, $4, $5);

-- name: SearchPerson :many
SELECT *  FROM Person
    WHERE to_tsvector('portuguese', nickname || ' ' || name || ' ' || COALESCE(f_stack_array(stack), ''))
          @@ plainto_tsquery(@query::text) LIMIT 50;
