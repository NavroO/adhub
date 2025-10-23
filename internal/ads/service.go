package ads

import (
	"context"
)

type Service interface {
	List(ctx context.Context) ([]Ad, error)
	Create(ctx context.Context, body CreateAdRequest) (Ad, error)
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

func (s *service) Create(ctx context.Context, body CreateAdRequest) (Ad, error) {
	return s.repo.Create(ctx, body)
}
