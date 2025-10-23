package ads

import (
	"context"
)

type Service interface {
	List(ctx context.Context) ([]Ad, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) List(ctx context.Context) ([]Ad, error) {
	return s.repo.List(ctx)
}
