package person

import (
	"context"
	"time"

	"github.com/avalonbits/round1fight/storage/pg/repo"
)

type Service struct {
	queries *repo.Queries
}

func New(queries *repo.Queries) *Service {
	return &Service{
		queries: queries,
	}
}

func (s *Service) Create(ctx context.Context, nickname, name string, birthday time.Time, stack []string) error {
	return nil
}
