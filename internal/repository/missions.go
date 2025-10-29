package repository

import (
	"context"
	"database/sql"

	"github.com/ilayzen/spy-cat-agency/pkg/models"
	"github.com/jmoiron/sqlx"
)

type missionRepository struct {
	ext sqlx.ExtContext
}

func NewMisionRepository(ext sqlx.ExtContext) *missionRepository {
	return &missionRepository{ext: ext}
}

func (ms *missionRepository) FetchMissions(ctx context.Context) ([]models.Mission, error) {
	query := `
		SELECT id, cat_id, completed, completed_at, created_at
		FROM missions
		ORDER BY id;
	`

	var missions []models.Mission
	if err := sqlx.SelectContext(ctx, ms.ext, &missions, query); err != nil {
		return nil, err
	}
	return missions, nil
}

func (ms *missionRepository) CreateMission(ctx context.Context, m models.Mission) (uint64, error) {
	query := `
		INSERT INTO missions (cat_id, completed, completed_at, created_at)
		VALUES ($1, $2, $3, NOW())
		RETURNING id;
	`
	var id uint64
	if err := sqlx.GetContext(ctx, ms.ext, &id, query, m.CatID, m.Completed, m.CompletedAt); err != nil {
		return 0, err
	}

	return id, nil
}

func (ms *missionRepository) GetMissionByID(ctx context.Context, id uint64) (models.Mission, error) {
	query := `
		SELECT id, cat_id, completed, completed_at, created_at
		FROM missions
		WHERE id = $1;
	`
	var mission models.Mission
	if err := sqlx.GetContext(ctx, ms.ext, &mission, query, id); err != nil {
		return models.Mission{}, err
	}
	return mission, nil
}

func (ms *missionRepository) DeleteMissionByID(ctx context.Context, id uint64) error {
	query := `DELETE FROM missions WHERE id = $1;`

	res, err := ms.ext.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

/*
	also we can use in the UpdateMissionByID next query:
		UPDATE missions
		SET completed = $1,
		    completed_at = CASE WHEN $1 = TRUE THEN NOW() ELSE NULL END
		WHERE id = $2;

	but I prefe use transactions for this cases

*/

func (ms *missionRepository) UpdateMissionByID(ctx context.Context, id uint64, m models.Mission) error {
	query := `
		UPDATE missions
		SET completed = $1,
		    completed_at = $2
		WHERE id = $3
	`
	res, err := ms.ext.ExecContext(ctx, query, m.Completed, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
