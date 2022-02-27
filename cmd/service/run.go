package service

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"github.com/penguin-statistics/backend-next/internal/config"
	"github.com/penguin-statistics/backend-next/internal/pkg/async"
)

func run(app *fiber.App, config *config.Config, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", config.Address)
			if err != nil {
				return err
			}

			go func() {
				if err := app.Listener(ln); err != nil {
					log.Error().Err(err).Msg("server terminated unexpectedly")
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if config.DevMode {
				return nil
			}

			return async.WaitAll(
				async.Errable(app.Shutdown),
				async.Errable(func() error {
					flushed := sentry.Flush(time.Second * 30)
					if !flushed {
						return errors.New("sentry flush timeout, some events may be lost")
					}
					return nil
				}),
			)
		},
	})
}
