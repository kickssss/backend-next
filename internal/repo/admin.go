package repo

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/penguin-statistics/backend-next/internal/models"
)

type Admin struct {
	db *bun.DB
}

func NewAdmin(db *bun.DB) *Admin {
	return &Admin{db: db}
}

func (r *Admin) SaveZones(ctx context.Context, tx bun.Tx, zones *[]*models.Zone) error {
	_, err := tx.NewInsert().
		On("CONFLICT (ark_zone_id) DO UPDATE").
		Model(zones).
		Exec(ctx)
	return err
}

func (r *Admin) SaveActivities(ctx context.Context, tx bun.Tx, activities *[]*models.Activity) error {
	_, err := tx.NewInsert().
		On("CONFLICT (activity_id) DO UPDATE").
		Model(activities).
		Exec(ctx)
	return err
}

func (r *Admin) SaveTimeRanges(ctx context.Context, tx bun.Tx, timeRanges *[]*models.TimeRange) error {
	_, err := tx.NewInsert().
		On("CONFLICT (range_id) DO UPDATE").
		Model(timeRanges).
		Exec(ctx)
	return err
}

func (r *Admin) SaveStages(ctx context.Context, tx bun.Tx, stages *[]*models.Stage) error {
	_, err := tx.NewInsert().
		On("CONFLICT (ark_stage_id) DO UPDATE").
		Model(stages).
		Exec(ctx)
	return err
}

func (r *Admin) SaveDropInfos(ctx context.Context, tx bun.Tx, dropInfos *[]*models.DropInfo) error {
	_, err := tx.NewInsert().
		On("CONFLICT (drop_id) DO UPDATE").
		Model(dropInfos).
		Exec(ctx)
	return err
}