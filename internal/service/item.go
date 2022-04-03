package service

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/ahmetb/go-linq/v3"
	"github.com/tidwall/gjson"

	"github.com/penguin-statistics/backend-next/internal/model"
	"github.com/penguin-statistics/backend-next/internal/model/cache"
	modelv2 "github.com/penguin-statistics/backend-next/internal/model/v2"
	"github.com/penguin-statistics/backend-next/internal/pkg/pgerr"
	"github.com/penguin-statistics/backend-next/internal/repo"
	"github.com/penguin-statistics/backend-next/internal/util"
)

type ItemService struct {
	ItemRepo *repo.Item
}

func NewItemService(itemRepo *repo.Item) *ItemService {
	return &ItemService{
		ItemRepo: itemRepo,
	}
}

// Cache: items, 24hrs
func (s *ItemService) GetItems(ctx context.Context) ([]*model.Item, error) {
	var items []*model.Item
	err := cache.Items.Get(&items)
	if err == nil {
		return items, nil
	}

	items, err = s.ItemRepo.GetItems(ctx)
	if err != nil {
		return nil, err
	}
	go cache.Items.Set(items, 24*time.Hour)
	return items, nil
}

func (s *ItemService) GetItemById(ctx context.Context, itemId int) (*model.Item, error) {
	itemsMapById, err := s.GetItemsMapById(ctx)
	if err != nil {
		return nil, err
	}
	item, ok := itemsMapById[itemId]
	if !ok {
		return nil, pgerr.ErrNotFound
	}
	return item, nil
}

// Cache: item#arkItemId:{arkItemId}, 24hrs
func (s *ItemService) GetItemByArkId(ctx context.Context, arkItemId string) (*model.Item, error) {
	var item model.Item
	err := cache.ItemByArkID.Get(arkItemId, &item)
	if err == nil {
		return &item, nil
	}

	dbItem, err := s.ItemRepo.GetItemByArkId(ctx, arkItemId)
	if err != nil {
		return nil, err
	}
	go cache.ItemByArkID.Set(arkItemId, *dbItem, 24*time.Hour)
	return dbItem, nil
}

func (s *ItemService) SearchItemByName(ctx context.Context, name string) (*model.Item, error) {
	return s.ItemRepo.SearchItemByName(ctx, name)
}

// Cache: (singular) shimItems, 24hrs; records last modified time
func (s *ItemService) GetShimItems(ctx context.Context) ([]*modelv2.Item, error) {
	var items []*modelv2.Item
	err := cache.ShimItems.Get(&items)
	if err == nil {
		return items, nil
	}

	items, err = s.ItemRepo.GetShimItems(ctx)
	if err != nil {
		return nil, err
	}
	for _, i := range items {
		s.applyShim(i)
	}
	if err := cache.ShimItems.Set(items, 24*time.Hour); err == nil {
		cache.LastModifiedTime.Set("[shimItems]", time.Now(), 0)
	}
	return items, nil
}

// Cache: shimItem#arkItemId:{arkItemId}, 24hrs
func (s *ItemService) GetShimItemByArkId(ctx context.Context, arkItemId string) (*modelv2.Item, error) {
	var item modelv2.Item
	err := cache.ShimItemByArkID.Get(arkItemId, &item)
	if err == nil {
		return &item, nil
	}

	dbItem, err := s.ItemRepo.GetShimItemByArkId(ctx, arkItemId)
	if err != nil {
		return nil, err
	}
	s.applyShim(dbItem)
	go cache.ShimItemByArkID.Set(arkItemId, *dbItem, 24*time.Hour)
	return dbItem, nil
}

// Cache: (singular) itemsMapById, 24hrs
func (s *ItemService) GetItemsMapById(ctx context.Context) (map[int]*model.Item, error) {
	var itemsMapById map[int]*model.Item
	cache.ItemsMapById.MutexGetSet(&itemsMapById, func() (map[int]*model.Item, error) {
		items, err := s.GetItems(ctx)
		if err != nil {
			return nil, err
		}
		s := make(map[int]*model.Item)
		for _, item := range items {
			s[item.ItemID] = item
		}
		return s, nil
	}, 24*time.Hour)
	return itemsMapById, nil
}

// Cache: (singular) itemsMapByArkId, 24hrs
func (s *ItemService) GetItemsMapByArkId(ctx context.Context) (map[string]*model.Item, error) {
	var itemsMapByArkId map[string]*model.Item
	cache.ItemsMapByArkID.MutexGetSet(&itemsMapByArkId, func() (map[string]*model.Item, error) {
		items, err := s.GetItems(ctx)
		if err != nil {
			return nil, err
		}
		s := make(map[string]*model.Item)
		for _, item := range items {
			s[item.ArkItemID] = item
		}
		return s, nil
	}, 24*time.Hour)
	return itemsMapByArkId, nil
}

func (s *ItemService) applyShim(item *modelv2.Item) {
	nameI18n := gjson.ParseBytes(item.NameI18n)
	item.Name = nameI18n.Map()["zh"].String()

	var coordSegments []int
	if item.Sprite.Valid {
		segments := strings.SplitN(item.Sprite.String, ":", 2)

		linq.From(segments).Select(func(i any) any {
			num, err := strconv.Atoi(i.(string))
			if err != nil {
				return -1
			}
			return num
		}).Where(func(i any) bool {
			return i.(int) != -1
		}).ToSlice(&coordSegments)
	}
	if coordSegments != nil {
		item.SpriteCoord = &coordSegments
	}

	keywords := gjson.ParseBytes(item.Keywords)

	item.AliasMap = json.RawMessage(util.Must(json.Marshal(keywords.Get("alias").Value())))
	item.PronMap = json.RawMessage(util.Must(json.Marshal(keywords.Get("pron").Value())))
}
