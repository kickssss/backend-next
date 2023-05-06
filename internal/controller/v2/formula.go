package v2

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"exusiai.dev/backend-next/internal/server/svr"
	"exusiai.dev/backend-next/internal/service"
)

type Formula struct {
	fx.In

	FormulaService *service.Formula
}

func RegisterFormula(v2 *svr.V2, c Formula) {
	v2.Get("/formula", c.GetFormula)
}

//	@Summary	Get Formula
//	@Tags		Formula
//	@Produce	json
//	@Success	200
//	@Failure	500	{object}	pgerr.PenguinError	"An unexpected error occurred"
//	@Router		/PenguinStats/api/v2/formula [GET]
func (c *Formula) GetFormula(ctx *fiber.Ctx) error {
	formula, err := c.FormulaService.GetFormula(ctx.UserContext())
	if err != nil {
		return err
	}
	return ctx.JSON(formula)
}
