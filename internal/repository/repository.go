package repository

import (
	"context"

	postgres "github.com/ilayzen/spy-cat-agency/pkg/database"
	"github.com/ilayzen/spy-cat-agency/pkg/models"
)

type Cats interface {
	Create(ctx context.Context, cat models.Cat) error
	Fetch(ctx context.Context) ([]models.Cat, error)
	GetByID(ctx context.Context, id uint64) (models.Cat, error)
	DeleteByID(ctx context.Context, id uint64) error
	UpdateByID(ctx context.Context, id uint64, salary uint64) error
}

type Missions interface {
	FetchMissions(ctx context.Context) ([]models.Mission, error)
	CreateMission(ctx context.Context, m models.Mission) (uint64, error)
	GetMissionByID(ctx context.Context, id uint64) (models.Mission, error)
	DeleteMissionByID(ctx context.Context, id uint64) error
	UpdateMissionByID(ctx context.Context, id uint64, m models.Mission) error
}

type Targets interface {
	AddMany(ctx context.Context, targets []models.Target) error
	UpdateByID(ctx context.Context, t models.Target) error
	DeleteByID(ctx context.Context, id uint64) error
	FetchByMissionID(ctx context.Context, missionID uint64) ([]models.Target, error)
	UpdateByMissionIDAndTargetID(ctx context.Context, missionID, targetID uint64, t models.Target) error
}

type Repository struct {
	*postgres.DB
	Cats
	Missions
	Targets
}

func NewRepository(db *postgres.DB) *Repository {
	return &Repository{
		db,
		NewCatsRepository(db),
		NewMisionRepository(db),
		NewTargetsRepository(db),
	}
}
