package cache

import (
	"sync"

	"github.com/go-redis/redis/v8"

	"github.com/penguin-statistics/backend-next/internal/utils/cache"
)

var (
	AccountById        cache.Cache
	AccountByPenguinId cache.Cache

	ItemDropSetByStageIdAndRangeId cache.Cache

	ShimMaxAccumulableDropMatrixResults cache.Cache

	Formula cache.Cache

	Items           cache.Cache
	ItemById        cache.Cache
	ItemByArkId     cache.Cache
	ShimItems       cache.Cache
	ShimItemByArkId cache.Cache

	Notices cache.Cache

	ShimLatestPatternMatrixResults cache.Cache

	ShimSiteStats cache.Cache

	once sync.Once
)

func Populate(client *redis.Client) {
	once.Do(func() {
		// account
		AccountById = cache.New(client, "account#accountId:")
		AccountByPenguinId = cache.New(client, "account#penguinId:")

		// drop_info
		ItemDropSetByStageIdAndRangeId = cache.New(client, "itemDropSet#server|stageId|rangeId:")

		// drop_matrix
		ShimMaxAccumulableDropMatrixResults = cache.New(client, "shimMaxAccumulableDropMatrixResults#server|showClosedZoned:")

		// formula
		Formula = cache.New(client, "formula")

		// item
		Items = cache.New(client, "items")
		ItemById = cache.New(client, "item#itemId:")
		ItemByArkId = cache.New(client, "item#arkItemId:")
		ShimItems = cache.New(client, "shimItems")
		ShimItemByArkId = cache.New(client, "shimItem#arkItemId:")

		// notice
		Notices = cache.New(client, "notices")

		// pattern_matrix
		ShimLatestPatternMatrixResults = cache.New(client, "shimLatestPatternMatrixResults#server:")

		// site_stats
		ShimSiteStats = cache.New(client, "shimSiteStats#server:")
	})
}
