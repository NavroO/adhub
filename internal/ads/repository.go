package ads

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	List(ctx context.Context) ([]Ad, error)
	Create(ctx context.Context, body CreateAdRequest) (Ad, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) List(ctx context.Context) ([]Ad, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, user_id, title, description, category, status, created_at, updated_at
		FROM ads
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}

	var ads []Ad
	for rows.Next() {
		var ad Ad
		if err := rows.Scan(
			&ad.ID, &ad.UserID, &ad.Title, &ad.Description,
			&ad.Category, &ad.Status, &ad.CreatedAt, &ad.UpdatedAt,
		); err != nil {
			return nil, err
		}
		ads = append(ads, ad)
	}
	return ads, nil
}

func (r *repository) Create(ctx context.Context, body CreateAdRequest) (Ad, error) {
	query := `
		INSERT INTO ads (user_id, title, description, category, status)
		VALUES ($1, $2, $3, $4, 'draft')
		RETURNING id, user_id, title, description, category, status, created_at, updated_at
	`

	var ad Ad
	err := r.db.QueryRowContext(
		ctx,
		query,
		body.UserID,
		body.Title,
		body.Description,
		body.Category,
	).Scan(
		&ad.ID,
		&ad.UserID,
		&ad.Title,
		&ad.Description,
		&ad.Category,
		&ad.Status,
		&ad.CreatedAt,
		&ad.UpdatedAt,
	)
	if err != nil {
		return Ad{}, fmt.Errorf("failed to create ad: %w", err)
	}

	return ad, nil
}
