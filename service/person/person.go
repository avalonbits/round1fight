package person

import (
	"context"
	"time"

	"github.com/avalonbits/round1fight/storage/pg/repo"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	queries *repo.Queries
}

func New(queries *repo.Queries) *Service {
	return &Service{
		queries: queries,
	}
}

func (s *Service) Create(
	ctx context.Context, nickname, name string, birthday time.Time, stack []string) (string, error) {
	id := uuid.New().String()
	return id, s.queries.CreatePerson(ctx, repo.CreatePersonParams{
		ID:       id,
		Nickname: nickname,
		Name:     name,
		Birthday: pgtype.Date{
			Time:  birthday,
			Valid: true,
		},
		Stack: stack,
	})
}
