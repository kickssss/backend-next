package repo

import (
	"context"
	"time"

	"github.com/uptrace/bun"

	"exusiai.dev/backend-next/internal/model"
	"exusiai.dev/backend-next/internal/repo/selector"
)

type PatternMatrixElement struct {
	db  *bun.DB
	sel selector.S[model.PatternMatrixElement]
}

func NewPatternMatrixElement(db *bun.DB) *PatternMatrixElement {
	return &PatternMatrixElement{db: db, sel: selector.New[model.PatternMatrixElement](db)}
}

func (s *PatternMatrixElement) BatchSaveElements(ctx context.Context, elements []*model.PatternMatrixElement, server string) error {
	_, err := s.db.NewInsert().Model(&elements).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *PatternMatrixElement) DeleteByServerAndDayNum(ctx context.Context, server string, dayNum int) error {
	_, err := s.db.NewDelete().Model((*model.PatternMatrixElement)(nil)).Where("server = ?", server).Where("day_num = ?", dayNum).Exec(ctx)
	return err
}

func (s *PatternMatrixElement) IsExistByServerAndDayNum(ctx context.Context, server string, dayNum int) (bool, error) {
	exists, err := s.db.NewSelect().Model((*model.PatternMatrixElement)(nil)).Where("server = ?", server).Where("day_num = ?", dayNum).Exists(ctx)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *PatternMatrixElement) GetAllTimesForGlobalPatternMatrix(
	ctx context.Context, server string, timeRange *model.TimeRange, stageIds []int, sourceCategory string,
) ([]*model.AllTimesResultForGlobalPatternMatrix, error) {
	subq2 := s.db.NewSelect().
		TableExpr("pattern_matrix_elements").
		Column("stage_id", "times", "day_num").
		Where("server = ?", server).
		Where("source_category = ?", sourceCategory).
		Where("times > 0").
		Where("stage_id IN (?)", bun.In(stageIds)).
		Where("start_time >= timestamp with time zone ?", timeRange.StartTime.Format(time.RFC3339)).
		Where("end_time <= timestamp with time zone ?", timeRange.EndTime.Format(time.RFC3339))

	subq1 := s.db.NewSelect().
		TableExpr("(?) AS subq2", subq2).
		Column("stage_id").
		ColumnExpr("MIN(times) AS times").
		Group("stage_id", "day_num")

	mainq := s.db.NewSelect().
		TableExpr("(?) AS subq1", subq1).
		Column("stage_id").
		ColumnExpr("SUM(times) AS times").
		Group("stage_id")

	results := make([]*model.AllTimesResultForGlobalPatternMatrix, 0)
	err := mainq.Scan(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *PatternMatrixElement) GetAllQuantitiesForGlobalPatternMatrix(
	ctx context.Context, server string, timeRange *model.TimeRange, stageIds []int, sourceCategory string,
) ([]*model.AllQuantitiesResultForGlobalPatternMatrix, error) {
	subq1 := s.db.NewSelect().
		TableExpr("pattern_matrix_elements").
		Column("stage_id", "pattern_id", "quantity").
		Where("server = ?", server).
		Where("source_category = ?", sourceCategory).
		Where("quantity > 0").
		Where("stage_id IN (?)", bun.In(stageIds)).
		Where("start_time >= timestamp with time zone ?", timeRange.StartTime.Format(time.RFC3339)).
		Where("end_time <= timestamp with time zone ?", timeRange.EndTime.Format(time.RFC3339))

	mainq := s.db.NewSelect().
		TableExpr("(?) AS subq1", subq1).
		Column("stage_id", "pattern_id").
		ColumnExpr("SUM(quantity) AS quantity").
		Group("stage_id", "pattern_id")

	results := make([]*model.AllQuantitiesResultForGlobalPatternMatrix, 0)
	err := mainq.Scan(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
