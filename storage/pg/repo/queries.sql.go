// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: queries.sql

package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countPerson = `-- name: CountPerson :one
SELECT COUNT(*) person_count from Person
`

func (q *Queries) CountPerson(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countPerson)
	var person_count int64
	err := row.Scan(&person_count)
	return person_count, err
}

const createPerson = `-- name: CreatePerson :exec
INSERT INTO Person (id, nickname, name, birthday, stack) VALUES ($1, $2, $3, $4, $5)
`

type CreatePersonParams struct {
	ID       string
	Nickname string
	Name     string
	Birthday pgtype.Date
	Stack    []string
}

func (q *Queries) CreatePerson(ctx context.Context, arg CreatePersonParams) error {
	_, err := q.db.Exec(ctx, createPerson,
		arg.ID,
		arg.Nickname,
		arg.Name,
		arg.Birthday,
		arg.Stack,
	)
	return err
}

const getPerson = `-- name: GetPerson :one
SELECT id, nickname, name, birthday, stack FROM Person where id = $1
`

func (q *Queries) GetPerson(ctx context.Context, id string) (Person, error) {
	row := q.db.QueryRow(ctx, getPerson, id)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.Nickname,
		&i.Name,
		&i.Birthday,
		&i.Stack,
	)
	return i, err
}

const searchPerson = `-- name: SearchPerson :many
SELECT id, nickname, name, birthday, stack  FROM Person WHERE person_fts_idx @@ to_tsquery($1::text) LIMIT 50
`

func (q *Queries) SearchPerson(ctx context.Context, query string) ([]Person, error) {
	rows, err := q.db.Query(ctx, searchPerson, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Person
	for rows.Next() {
		var i Person
		if err := rows.Scan(
			&i.ID,
			&i.Nickname,
			&i.Name,
			&i.Birthday,
			&i.Stack,
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
