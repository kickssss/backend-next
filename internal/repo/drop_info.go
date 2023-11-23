package repo

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"exusiai.dev/gommon/constant"
	"github.com/ahmetb/go-linq/v3"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/uptrace/bun"

	"exusiai.dev/backend-next/internal/model"
	"exusiai.dev/backend-next/internal/pkg/pgerr"
	"exusiai.dev/backend-next/internal/pkg/pgqry"
	"exusiai.dev/backend-next/internal/repo/selector"
)

type DropInfo struct {
	db  *bun.DB
	sel selector.S[model.DropInfo]
}

func NewDropInfo(db *bun.DB) *DropInfo {
	return &DropInfo{db: db, sel: selector.New[model.DropInfo](db)}
}

func (r *DropInfo) GetDropInfo(ctx context.Context, id int) (*model.DropInfo, error) {
	return r.sel.SelectOne(ctx, func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("id = ?", id)
	})
}

func (r *DropInfo) GetDropInfosByServerAndStageId(ctx context.Context, server string, stageId int) ([]*model.DropInfo, error) {
	return r.sel.SelectMany(ctx, func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("stage_id = ?", stageId).Where("server = ?", server)
	})
}

func (r *DropInfo) GetDropInfosByServer(ctx context.Context, server string) ([]*model.DropInfo, error) {
	return r.sel.SelectMany(ctx, func(q *bun.SelectQuery) *bun.SelectQuery {
		return q.Where("server = ?", server)
	})
}

type DropInfoQuery struct {
	Server     string
	ArkStageId string
}

// GetDropInfoByArkId returns a drop info by its ark id.
func (r *DropInfo) GetForCurrentTimeRange(ctx context.Context, query *DropInfoQuery) ([]*model.DropInfo, error) {
	var dropInfo []*model.DropInfo
	err := pgqry.New(
		r.db.NewSelect().
			Model(&dropInfo).
			Where("di.server = ?", query.Server).
			Where("st.ark_stage_id = ?", query.ArkStageId),
	).
		UseItemById("di.item_id").
		UseStageById("di.stage_id").
		UseTimeRange("di.range_id").
		DoFilterCurrentTimeRange().
		Q.Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pgerr.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return dropInfo, nil
}

func (r *DropInfo) GetItemDropSetByStageIdAndRangeId(ctx context.Context, server string, stageId int, rangeId int) ([]int, error) {
	var results []int
	err := r.db.NewSelect().
		Column("di.item_id").
		Model((*model.DropInfo)(nil)).
		Where("di.server = ?", server).
		Where("di.stage_id = ?", stageId).
		Where("di.item_id IS NOT NULL").
		Where("di.range_id = ?", rangeId).
		Scan(ctx, &results)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pgerr.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	results = lo.Uniq(results)
	sort.Ints(results)
	return results, nil
}

func (r *DropInfo) GetForCurrentTimeRangeWithDropTypes(ctx context.Context, query *DropInfoQuery) (itemDropInfos, typeDropInfos []*model.DropInfo, err error) {
	allDropInfos, err := r.GetForCurrentTimeRange(ctx, query)
	if err != nil {
		return nil, nil, err
	}
	for _, dropInfo := range allDropInfos {
		if dropInfo.DropType != constant.DropTypeRecognitionOnly {
			if dropInfo.ItemID.Valid {
				itemDropInfos = append(itemDropInfos, dropInfo)
			} else {
				typeDropInfos = append(typeDropInfos, dropInfo)
			}
		}
	}

	return itemDropInfos, typeDropInfos, nil
}

func (r *DropInfo) GetDropInfosWithFilters(ctx context.Context, server string, timeRanges []*model.TimeRange, stageIdFilter []int, itemIdFilter []int) ([]*model.DropInfo, error) {
	results := make([]*model.DropInfo, 0)
	var whereBuilder strings.Builder
	fmt.Fprintf(&whereBuilder, "di.server = ? AND di.drop_type != ? AND di.item_id IS NOT NULL")

	if len(stageIdFilter) > 0 {
		fmt.Fprintf(&whereBuilder, " AND di.stage_id IN (?)")
	}
	if len(itemIdFilter) > 0 {
		fmt.Fprintf(&whereBuilder, " AND di.item_id IN (?)")
	}

	allTimeRangesHaveNoRangeId := true
	for _, timeRange := range timeRanges {
		if timeRange.RangeID > 0 {
			allTimeRangesHaveNoRangeId = false
			break
		}
	}

	if len(timeRanges) > 0 {
		if allTimeRangesHaveNoRangeId {
			for _, timeRange := range timeRanges {
				startTimeStr := timeRange.StartTime.Format(time.RFC3339)
				endTimeStr := timeRange.EndTime.Format(time.RFC3339)
				fmt.Fprintf(&whereBuilder,
					" AND tr.start_time <= timestamp with time zone '%s' AND tr.end_time >= timestamp with time zone '%s'",
					endTimeStr,
					startTimeStr)
			}
		} else {
			if len(timeRanges) == 1 {
				fmt.Fprintf(&whereBuilder, " AND di.range_id = %d", timeRanges[0].RangeID)
			} else {
				rangeIdStr := make([]string, len(timeRanges))
				linq.From(timeRanges).SelectT(func(timeRange *model.TimeRange) string { return strconv.Itoa(timeRange.RangeID) }).ToSlice(&rangeIdStr)
				fmt.Fprintf(&whereBuilder, " AND di.range_id IN (%s)", strings.Join(rangeIdStr, ","))
			}
		}
	}
	if err := r.db.NewSelect().TableExpr("drop_infos as di").Column("di.stage_id", "di.item_id", "di.accumulable").
		Where(whereBuilder.String(), server, constant.DropTypeRecognitionOnly, bun.In(stageIdFilter), bun.In(itemIdFilter)).
		Join("JOIN time_ranges AS tr ON tr.range_id = di.range_id").
		Scan(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (s *DropInfo) GetDropInfosByServerAndRangeId(ctx context.Context, server string, rangeId int) ([]*model.DropInfo, error) {
	var dropInfo []*model.DropInfo
	err := s.DB.NewSelect().
		Model(&dropInfo).
		Where("server = ?", server).
		Where("range_id = ?", rangeId).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pgerr.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return dropInfo, nil
}

func (s *DropInfo) CloneDropInfosFromCN(ctx context.Context, originRangeId int, destRangeId int, server string) error {
	dropInfos, err := s.GetDropInfosByServerAndRangeId(ctx, "CN", originRangeId)
	if err != nil {
		return err
	}
	for _, dropInfo := range dropInfos {
		dropInfo.DropID = 0
		dropInfo.RangeID = destRangeId
		dropInfo.Server = server
	}
	_, err = s.DB.NewInsert().Model(&dropInfos).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
