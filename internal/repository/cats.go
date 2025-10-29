package repository

import (
	"context"
	"database/sql"

	"github.com/ilayzen/spy-cat-agency/pkg/models"
	"github.com/jmoiron/sqlx"
)

type CatsRepository struct {
	ext sqlx.ExtContext
}

func NewCatsRepository(ext sqlx.ExtContext) *CatsRepository {
	return &CatsRepository{
		ext: ext,
	}
}

func (cs *CatsRepository) Create(ctx context.Context, cat models.Cat) error {
	query := `
		INSERT INTO cats (name, years_experience, breed, salary)
		VALUES ($1, $2, $3, $4);
	`

	_, err := cs.ext.ExecContext(ctx, query,
		cat.Name,
		cat.YearsExperience,
		cat.Breed,
		cat.Salary,
	)
	return err
}

func (cs *CatsRepository) Fetch(ctx context.Context) ([]models.Cat, error) {
	query := `SELECT * FROM cats`

	var cats []models.Cat
	err := sqlx.SelectContext(ctx, cs.ext, &cats, query)
	if err != nil {
		return nil, err
	}

	return cats, nil
}

func (r *CatsRepository) GetByID(ctx context.Context, id uint64) (models.Cat, error) {
	query := `
		SELECT id, name, years_experience, breed, salary, created_at
		FROM cats
		WHERE id = $1;
	`

	var cat models.Cat
	err := sqlx.GetContext(ctx, r.ext, &cat, query, id)
	if err != nil {
		return models.Cat{}, err
	}

	return cat, nil
}

func (r *CatsRepository) DeleteByID(ctx context.Context, id uint64) error {
	query := `DELETE FROM cats WHERE id = $1;`

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

func (r *CatsRepository) UpdateByID(ctx context.Context, id uint64, salary uint64) error {
	query := `
		UPDATE cats
		SET salary = $1
		WHERE id = $2;
	`

	result, err := r.ext.ExecContext(ctx, query, salary, id)
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
