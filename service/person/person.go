package person

import (
	"context"
	"fmt"
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

func (s *Service) Count(ctx context.Context) (int64, error) {
	return s.queries.CountPerson(ctx)
}

type SearchResult struct {
}

func (s *Service) Search(ctx context.Context, query string) ([]SearchResult, error) {
	res, err := s.queries.SearchPerson(ctx, query)
	if err != nil {
		return nil, err
	}

	var results []SearchResult
	for _, r := range res {
		fmt.Printf("%#v+\n", r)
		results = append(results, SearchResult{})
	}
	return results, nil
}
