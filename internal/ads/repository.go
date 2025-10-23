package ads

import (
	"context"
	"database/sql"
)

type Repository interface {
	List(ctx context.Context) ([]Ad, error)
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
