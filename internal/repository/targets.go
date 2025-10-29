package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/ilayzen/spy-cat-agency/pkg/models"
	"github.com/jmoiron/sqlx"
)

type targetsRepository struct {
	ext sqlx.ExtContext
}

func NewTargetsRepository(ext sqlx.ExtContext) *targetsRepository {
	return &targetsRepository{ext: ext}
}

func (r *targetsRepository) AddMany(ctx context.Context, ts []models.Target) error {
	if len(ts) == 0 {
		return nil
	}

	var sb strings.Builder
	args := make([]any, 0, len(ts)*6)

	sb.WriteString(`INSERT INTO targets (mission_id, name, country, notes, completed, completed_at) VALUES `)
	idx := 1
	for i, t := range ts {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d)", idx, idx+1, idx+2, idx+3, idx+4, idx+5))
		args = append(args, t.MissionID, t.Name, t.Country, nullableString(t.Notes), t.Completed, t.CompletedAt)
		idx += 6
	}
	_, err := r.ext.ExecContext(ctx, sb.String(), args...)
	return err
}

func nullableString(s string) any {
	if s == "" {
		return nil
	}
	return s
}

func (r *targetsRepository) UpdateByID(ctx context.Context, t models.Target) error {
	query := `
		UPDATE targets
		SET 
			notes = COALESCE($1, notes),
			completed = COALESCE($2, completed),
			completed_at = CASE 
				WHEN $2 = TRUE THEN NOW() 
				WHEN $2 = FALSE THEN NULL
				ELSE completed_at 
			END,
			updated_at = NOW()
		WHERE id = $3;
	`

	result, err := r.ext.ExecContext(ctx, query,
		t.Notes,
		t.Completed,
		t.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *targetsRepository) DeleteByID(ctx context.Context, id uint64) error {
	query := `DELETE FROM targets WHERE id = $1;`

	result, err := r.ext.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *targetsRepository) FetchByMissionID(ctx context.Context, missionID uint64) ([]models.Target, error) {
	query := `
		SELECT 
			id, mission_id, name, country, notes, 
			completed, completed_at, created_at
		FROM targets
		WHERE mission_id = $1
		ORDER BY id;
	`

	var targets []models.Target
	if err := sqlx.SelectContext(ctx, r.ext, &targets, query, missionID); err != nil {
		return nil, err
	}

	return targets, nil
}

func (r *targetsRepository) UpdateByMissionIDAndTargetID(
	ctx context.Context,
	missionID, targetID uint64,
	t models.Target,
) error {
	const q = `
        UPDATE targets
           SET name         = $1,
               country      = $2,
               notes        = $3,
               completed    = $4,
               completed_at = $5
         WHERE id = $6
           AND mission_id = $7;
    `

	res, err := r.ext.ExecContext(ctx, q,
		t.Name,
		t.Country,
		nullableString(t.Notes),
		t.Completed,
		t.CompletedAt,
		targetID,
		missionID,
	)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err == nil && n == 0 {
		return sql.ErrNoRows
	}
	return err
}
