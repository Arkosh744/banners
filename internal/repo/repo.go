package repo

import (
	"context"
	"fmt"

	"github.com/Arkosh744/banners/pkg/models"
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
		return 0, fmt.Errorf("failed to create slot: %w", err)
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
		return 0, fmt.Errorf("failed to create banner: %w", err)
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
		return 0, fmt.Errorf("failed to create soc group: %w", err)
	}

	return id, nil
}

func (r *Repo) AddBannerToSlot(ctx context.Context, req *models.BannerSlotRequest) error {
	query := `INSERT INTO banner_slot (slot_id, banner_id) VALUES ($1, $2);`

	q := pg.Query{
		Name:     "banners.AddBannerToSlot",
		QueryRaw: query,
	}

	if _, err := r.db.PG().ExecContext(ctx, q, req.SlotID, req.BannerID); err != nil {
		return fmt.Errorf("failed to add banner to slot: %w", err)
	}

	return nil
}

func (r *Repo) DeleteBannerSlot(ctx context.Context, req *models.BannerSlotRequest) error {
	query := `DELETE FROM banner_slot WHERE slot_id = $1 AND banner_id = $2;`

	q := pg.Query{
		Name:     "banners.DeleteBannerSlot",
		QueryRaw: query,
	}

	if _, err := r.db.PG().ExecContext(ctx, q, req.SlotID, req.BannerID); err != nil {
		return fmt.Errorf("failed to delete banner from slot: %w", err)
	}

	return nil
}

func (r *Repo) CreateClickEvent(ctx context.Context, req *models.EventRequest) error {
	query := `INSERT INTO clicks (slot_id, banner_id, group_id) VALUES ($1, $2, $3);`

	q := pg.Query{
		Name:     "banners.CreateClickEvent",
		QueryRaw: query,
	}

	if _, err := r.db.PG().ExecContext(ctx, q, req.SlotID, req.BannerID, req.GroupID); err != nil {
		return fmt.Errorf("failed to create click event: %w", err)
	}

	return nil
}

func (r *Repo) GetBannersInfo(ctx context.Context, req *models.NextBannerRequest) ([]models.BannerStats, error) {
	query := `SELECT bs.banner_id,
					bs.slot_id, 
					bv.group_id, 
				    count(distinct bv.id) view_count, 
					count(distinct cl.id) click_count
			FROM banner_slot bs
			LEFT JOIN views bv ON bv.slot_id = bs.slot_id AND bv.banner_id = bs.banner_id
			LEFT JOIN clicks cl ON bv.slot_id = cl.slot_id AND bv.banner_id = cl.banner_id AND 
											   bv.group_id = cl.group_id
			WHERE bs.slot_id = $1 AND (bv.group_id = $2 OR bv.group_id is null)
			GROUP BY bs.banner_id, bs.slot_id, bv.group_id
			ORDER BY bv.group_id`

	q := pg.Query{
		Name:     "banners.GetBannersInfo",
		QueryRaw: query,
	}

	var banners []models.BannerStats
	if err := r.db.PG().ScanAllContext(ctx, &banners, q, req.SlotID, req.GroupID); err != nil {
		return nil, fmt.Errorf("failed to get banners info: %w", err)
	}

	for i := range banners {
		if banners[i].GroupID == nil {
			banners[i].GroupID = &req.GroupID
		}
	}

	return banners, nil
}

func (r *Repo) IncrementBannerView(ctx context.Context, req *models.EventRequest) error {
	query := `INSERT INTO views (banner_id, slot_id, group_id) VALUES ($1, $2, $3)`

	q := pg.Query{
		Name:     "banners.IncrementBannerView",
		QueryRaw: query,
	}

	if _, err := r.db.PG().ExecContext(ctx, q, req.BannerID, req.SlotID, req.GroupID); err != nil {
		return fmt.Errorf("failed to increment banner view: %w", err)
	}

	return nil
}
