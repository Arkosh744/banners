package repo

import (
	"context"
	"github.com/Arkosh744/banners/pkg/pg"
)

type Repo struct {
	db pg.Client
}

func NewRepo(client pg.Client) *Repo {
	return &Repo{db: client}
}

func (r *Repo) CreateSlot(ctx context.Context, description string) (int, error) {
	query := `INSERT INTO slots (description) VALUES ($1) RETURNING id;`

	q := pg.Query{
		Name:     "banners.CreateSlot",
		QueryRaw: query,
	}

	var id int
	if err := r.db.PG().ScanOneContext(ctx, &id, q, description); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repo) CreateBanner(ctx context.Context, description string) (int, error) {
	query := `INSERT INTO banners (description) VALUES ($1) RETURNING id;`

	q := pg.Query{
		Name:     "banners.CreateBanner",
		QueryRaw: query,
	}

	var id int
	if err := r.db.PG().ScanOneContext(ctx, &id, q, description); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repo) CreateSocGroup(ctx context.Context, description string) (int, error) {
	query := `INSERT INTO social_groups (description) VALUES ($1) RETURNING id;`

	q := pg.Query{
		Name:     "banners.CreateSocGroup",
		QueryRaw: query,
	}

	var id int
	if err := r.db.PG().ScanOneContext(ctx, &id, q, description); err != nil {
		return 0, err
	}

	return id, nil
}
