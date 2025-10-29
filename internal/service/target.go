package service

import (
	"context"

	"github.com/ilayzen/spy-cat-agency/internal/repository"
	"github.com/ilayzen/spy-cat-agency/pkg/models"
	"github.com/sirupsen/logrus"
)

type targetService struct {
	log  logrus.FieldLogger
	repo *repository.Repository
}

func NewTargetService(log logrus.FieldLogger, repo *repository.Repository) *targetService {
	return &targetService{
		log:  log,
		repo: repo,
	}
}

func (cs *targetService) UpdateByMissionIDAndTargetID(ctx context.Context, missionID, targetID uint64, target models.Target) error {
	return cs.repo.Targets.UpdateByMissionIDAndTargetID(ctx, missionID, targetID, target)
}
