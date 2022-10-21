package appentry

import (
	"time"

	"go.uber.org/fx"

	"exusiai.dev/backend-next/internal/config"
	controllermeta "exusiai.dev/backend-next/internal/controller/meta"
	controllerv2 "exusiai.dev/backend-next/internal/controller/v2"
	controllerv3 "exusiai.dev/backend-next/internal/controller/v3"
	"exusiai.dev/backend-next/internal/infra"
	"exusiai.dev/backend-next/internal/model/cache"
	"exusiai.dev/backend-next/internal/pkg/crypto"
	"exusiai.dev/backend-next/internal/pkg/logger"
	"exusiai.dev/backend-next/internal/repo"
	"exusiai.dev/backend-next/internal/server/httpserver"
	"exusiai.dev/backend-next/internal/server/svr"
	"exusiai.dev/backend-next/internal/service"
	"exusiai.dev/backend-next/internal/util/reportverifs"
	"exusiai.dev/backend-next/internal/workers/calcwkr"
	"exusiai.dev/backend-next/internal/workers/reportwkr"
)

func ProvideOptions(includeSwagger bool) []fx.Option {
	opts := []fx.Option{
		// Misc
		fx.Provide(config.Parse),
		fx.Provide(httpserver.Create),
		fx.Provide(svr.CreateEndpointGroups),
		fx.Provide(crypto.NewCrypto),

		// Infrastructures
		infra.Module(),

		// Verifiers
		reportverifs.Module(),

		// Repositories
		repo.Module(),

		// Services
		service.Module(),

		// Global Singleton Inits: Keep those before controllers to ensure they are initialized
		// before controllers are registered as controllers are also fx#Invoke functions which
		// are called in the order of their registration.
		fx.Invoke(logger.Configure),
		fx.Invoke(infra.SentryInit),
		fx.Invoke(cache.Initialize),
		fx.WithLogger(logger.Fx),

		// Controllers (v2)
		controllerv2.Module(),

		// Controllers (v3)
		controllerv3.Module(),

		// Controllers (meta)
		controllermeta.Module(),

		// Workers
		fx.Invoke(calcwkr.Start),
		fx.Invoke(reportwkr.Start),

		// fx Extra Options
		fx.StartTimeout(1 * time.Second),
		// StopTimeout is not typically needed, since we're using fiber's Shutdown(),
		// in which fiber has its own IdleTimeout for controlling the shutdown timeout.
		// It acts as a countermeasure in case the fiber app is not properly shutting down.
		fx.StopTimeout(5 * time.Minute),
	}

	if includeSwagger {
		opts = append(opts, fx.Invoke(controllermeta.RegisterSwagger))
	}

	return opts
}
