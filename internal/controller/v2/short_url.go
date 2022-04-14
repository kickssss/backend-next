package v2

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"github.com/penguin-statistics/backend-next/internal/core/shorturl"
	"github.com/penguin-statistics/backend-next/internal/server/svr"
)

type ShortURL struct {
	fx.In

	ShortURLService *shorturl.Service
}

func RegisterShortURL(v2 *svr.V2, c ShortURL) {
	v2.Get("/short", c.Resolve)
	v2.Get("/short/:word", c.Resolve)
}

func (c *ShortURL) Resolve(ctx *fiber.Ctx) error {
	word := ctx.Params("word")
	return ctx.Redirect(c.ShortURLService.Resolve(ctx, word))
}
