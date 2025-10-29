package service

import (
	"context"

	"github.com/ilayzen/spy-cat-agency/internal/repository"
	"github.com/ilayzen/spy-cat-agency/pkg/models"
	"github.com/sirupsen/logrus"
)

type catsService struct {
	log  logrus.FieldLogger
	repo *repository.Repository
}

func NewCatsService(log logrus.FieldLogger, repo *repository.Repository) *catsService {
	return &catsService{
		log:  log,
		repo: repo,
	}
}

func (cs *catsService) Create(ctx context.Context, cat models.Cat) error {
	return cs.repo.Cats.Create(ctx, cat)
}

func (cs *catsService) Fetch(ctx context.Context) ([]models.Cat, error) {
	return cs.repo.Cats.Fetch(ctx)
}

func (cs *catsService) GetByID(ctx context.Context, id uint64) (models.Cat, error) {
	return cs.repo.Cats.GetByID(ctx, id)
}

func (cs *catsService) DeleteByID(ctx context.Context, id uint64) error {
	return cs.repo.Cats.DeleteByID(ctx, id)
}

func (cs *catsService) UpdateByID(ctx context.Context, id uint64, salary uint64) error {
	return cs.repo.Cats.UpdateByID(ctx, id, salary)
}
