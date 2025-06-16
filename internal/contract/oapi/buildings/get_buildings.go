package buildings

import (
	"OMS_assignment/internal/contract/oapi/buildings/dto"
	"github.com/gofiber/fiber/v2"
)

func (bc *Contract) GetBuildings(c *fiber.Ctx) error {
	ctx := c.Context()

	buildings, err := bc.buildingsUsecase.GetBuildings(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			errorField: err.Error(),
		})
	}

	resp := dto.BuildingsToResponse(buildings)
	return c.Status(fiber.StatusOK).JSON(resp)
}
