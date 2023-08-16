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

func (s *Service) Count(ctx context.Context) (int64, error) {
	return s.queries.CountPerson(ctx)
}

type SearchResult struct {
	ID       string   `json:"id"`
	Nickname string   `json:"apelido"`
	Name     string   `json:"name"`
	Birthday string   `json:"nascimento"`
	Stack    []string `json:"stack"`
}

func (s *Service) Search(ctx context.Context, query string) ([]SearchResult, error) {
	res, err := s.queries.SearchPerson(ctx, query)
	if err != nil {
		return nil, err
	}

	results := []SearchResult{}
	for _, r := range res {
		results = append(results, SearchResult{
			ID:       r.ID,
			Nickname: r.Nickname,
			Name:     r.Name,
			Stack:    r.Stack,
			Birthday: r.Birthday.Time.Format("2006-01-02"),
		})
	}
	return results, nil
}

func (s *Service) Get(ctx context.Context, id string) (SearchResult, error) {
	r, err := s.queries.GetPerson(ctx, id)
	if err != nil {
		return SearchResult{}, err
	}
	return SearchResult{
		ID:       r.ID,
		Nickname: r.Nickname,
		Name:     r.Name,
		Stack:    r.Stack,
		Birthday: r.Birthday.Time.Format("2006-01-02"),
	}, nil
}
