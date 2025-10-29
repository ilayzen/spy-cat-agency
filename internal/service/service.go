package service

import (
	"context"

	"github.com/ilayzen/spy-cat-agency/internal/repository"
	"github.com/ilayzen/spy-cat-agency/pkg/models"
	"github.com/sirupsen/logrus"
)

type Cats interface {
	Create(ctx context.Context, cat models.Cat) error
	Fetch(ctx context.Context) ([]models.Cat, error)
	GetByID(ctx context.Context, id uint64) (models.Cat, error)
	DeleteByID(ctx context.Context, id uint64) error
	UpdateByID(ctx context.Context, id uint64, salary uint64) error
}

type Missions interface {
	FetchMissions(ctx context.Context) ([]models.ResponseMission, error)
	CreateMission(ctx context.Context, m models.RequestMission) error
	GetMissionByID(ctx context.Context, id uint64) (models.Mission, error)
	DeleteMissionByID(ctx context.Context, id uint64) error
	UpdateMissionByID(ctx context.Context, id uint64, m models.Mission) error
}

type Target interface {
	UpdateByMissionIDAndTargetID(ctx context.Context, missionID, targetID uint64, target models.Target) error
}

type Service struct {
	Cats
	Missions
	Target
}

func NewService(log logrus.FieldLogger, repos *repository.Repository) *Service {
	return &Service{
		Cats:     NewCatsService(log, repos),
		Missions: NewMissionsService(log, repos),
		Target:   NewTargetService(log, repos),
	}
}
