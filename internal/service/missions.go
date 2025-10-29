package service

import (
	"context"
	"fmt"

	"github.com/ilayzen/spy-cat-agency/internal/repository"
	"github.com/ilayzen/spy-cat-agency/pkg/models"
	"github.com/sirupsen/logrus"
)

type missionsService struct {
	log  logrus.FieldLogger
	repo *repository.Repository
}

func NewMissionsService(log logrus.FieldLogger, repo *repository.Repository) *missionsService {
	return &missionsService{
		log:  log,
		repo: repo,
	}
}

func (cs *missionsService) FetchMissions(ctx context.Context) ([]models.ResponseMission, error) {
	missions, err := cs.repo.Missions.FetchMissions(ctx)
	if err != nil {
		return nil, err
	}

	var resp []models.ResponseMission

	catRepo := repository.NewCatsRepository(cs.repo.DB)
	targetRepo := repository.NewTargetsRepository(cs.repo.DB)

	for _, m := range missions {
		r := models.ResponseMission{
			Mission: m,
		}

		if m.CatID != nil {
			cat, err := catRepo.GetByID(ctx, uint64(*m.CatID))
			if err != nil {
				return nil, err
			}
			r.Cat = cat
		}

		targets, err := targetRepo.FetchByMissionID(ctx, uint64(m.ID))
		if err != nil {
			return nil, err
		}
		r.Targets = targets

		resp = append(resp, r)
	}

	return resp, nil
}

func (cs *missionsService) CreateMission(ctx context.Context, m models.RequestMission) (err error) {
	tx, err := cs.repo.Beginx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	missionRepo := repository.NewMisionRepository(tx)

	id, err := missionRepo.CreateMission(ctx, m.Mission)
	if err != nil {
		return err
	}

	if len(m.Targets) > 0 {
		for i := range m.Targets {
			m.Targets[i].MissionID = int64(id)
		}
		targetRepo := repository.NewTargetsRepository(tx)
		if err = targetRepo.AddMany(ctx, m.Targets); err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (cs *missionsService) GetMissionByID(ctx context.Context, id uint64) (models.Mission, error) {
	return cs.repo.Missions.GetMissionByID(ctx, id)
}

func (cs *missionsService) DeleteMissionByID(ctx context.Context, id uint64) error {
	m, err := cs.repo.Missions.GetMissionByID(ctx, id)
	if err != nil {
		return nil
	}

	if m.CatID == nil {
		return fmt.Errorf("cannot delete mission, cat was assignet to the mission")
	}

	return cs.repo.Missions.DeleteMissionByID(ctx, id)
}

func (cs *missionsService) UpdateMissionByID(ctx context.Context, id uint64, m models.Mission) error {
	return cs.repo.Missions.UpdateMissionByID(ctx, id, m)
}
