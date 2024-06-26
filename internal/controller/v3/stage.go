package v3

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"exusiai.dev/backend-next/internal/server/svr"
	"exusiai.dev/backend-next/internal/service"
)

type StageController struct {
	fx.In

	StageService *service.Stage
}

func RegisterStage(v3 *svr.V3, c StageController) {
	v3.Get("/stages", c.GetStages)
	v3.Get("/stages/:stageId", c.GetStageById)
}

func (c *StageController) GetStages(ctx *fiber.Ctx) error {
	stages, err := c.StageService.GetStages(ctx.UserContext())
	if err != nil {
		return err
	}

	return ctx.JSON(stages)
}

func (c *StageController) GetStageById(ctx *fiber.Ctx) error {
	stageId := ctx.Params("stageId")

	stage, err := c.StageService.GetStageByArkId(ctx.UserContext(), stageId)
	if err != nil {
		return err
	}

	return ctx.JSON(stage)
}
