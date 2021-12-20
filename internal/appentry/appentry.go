package appentry

import (
	"time"

	"github.com/penguin-statistics/backend-next/internal/config"
	"github.com/penguin-statistics/backend-next/internal/controllers"
	"github.com/penguin-statistics/backend-next/internal/controllers/shims"
	"github.com/penguin-statistics/backend-next/internal/infra"
	"github.com/penguin-statistics/backend-next/internal/pkg/flake"
	"github.com/penguin-statistics/backend-next/internal/repos"
	"github.com/penguin-statistics/backend-next/internal/server"
	httpserver "github.com/penguin-statistics/backend-next/internal/server/http"
	"github.com/penguin-statistics/backend-next/internal/utils"

	"go.uber.org/fx"
)

func ProvideOptions(includeSwagger bool) []fx.Option {
	opts := []fx.Option{
		fx.Invoke(utils.ConfigureLogger),
		fx.Provide(config.Parse),
		fx.Provide(flake.NewSnowflake),
		fx.Provide(httpserver.CreateServer),
		fx.Provide(infra.ProvideNats),
		fx.Provide(infra.ProvidePostgres),
		fx.Provide(infra.ProvideRedis),
		fx.Provide(repos.NewItemRepo),
		fx.Provide(repos.NewStageRepo),
		fx.Provide(repos.NewZoneRepo),
		fx.Provide(server.CreateVersioningEndpoints),
		fx.Provide(utils.NewCrypto),
		fx.Invoke(shims.RegisterItemController),
		fx.Invoke(shims.RegisterStageController),
		fx.Invoke(shims.RegisterZoneController),
		fx.Invoke(controllers.RegisterIndexController),
		fx.Invoke(controllers.RegisterItemController),
		fx.Invoke(controllers.RegisterStageController),
		fx.Invoke(controllers.RegisterZoneController),
		fx.Invoke(controllers.RegisterReportController),
		fx.Invoke(controllers.RegisterLiveController),
		fx.StartTimeout(1 * time.Second),
		// StopTimeout is not typically needed, since we're using fiber's Shutdown(),
		// in which fiber has its own IdleTimeout for controlling the shutdown timeout.
		// It acts as a countermeasure in case the fiber app is not properly shutting down.
		fx.StopTimeout(5 * time.Minute),
	}

	if includeSwagger {
		opts = append(opts, fx.Invoke(controllers.RegisterSwaggerController))
	}

	return opts
}
