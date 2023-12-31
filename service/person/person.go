package person

import (
	"context"
	"time"

	"github.com/avalonbits/round1fight/storage/pg/repo"
	"github.com/google/uuid"
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
		Birthday: birthday.Format("2006-01-02"),
		Stack:    stack,
	})
}

func (s *Service) Count(ctx context.Context) (int64, error) {
	return s.queries.CountPerson(ctx)
}

//easyjson:json
type Result struct {
	ID       string   `json:"id"`
	Nickname string   `json:"apelido"`
	Name     string   `json:"nome"`
	Birthday string   `json:"nascimento"`
	Stack    []string `json:"stack"`
}

//easyjson:json
type Results []Result

func (s *Service) Search(ctx context.Context, query string) (Results, error) {
	res, err := s.queries.SearchPerson(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return Results{}, nil
	}

	results := make(Results, 0, len(res))
	for _, r := range res {
		results = append(results, resultFrom(r))
	}
	return results, nil
}

func (s *Service) Get(ctx context.Context, id string) (Result, error) {
	r, err := s.queries.GetPerson(ctx, id)
	if err != nil {
		return Result{}, err
	}
	return resultFrom(repo.SearchPersonRow(r)), nil
}

func resultFrom(r repo.SearchPersonRow) Result {
	return Result{
		ID:       r.ID,
		Nickname: r.Nickname,
		Name:     r.Name,
		Stack:    r.Stack,
		Birthday: r.Birthday,
	}
}
