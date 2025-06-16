package apartments

import (
	"OMS_assignment/internal/contract/oapi/apartments/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (ac *Contract) GetApartmentsByBuildingID(c *fiber.Ctx) error {
	buildingIDRaw := c.Params(BuildingIDParam)

	buildingID, err := strconv.ParseInt(buildingIDRaw, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			errorField: err.Error(),
		})
	}
	ctx := c.Context()
	apartments, err := ac.apartmentsUsecase.GetApartmentsByBuildingID(ctx, buildingID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			errorField: err.Error(),
		})
	}

	resp := dto.ApartmentsToResponse(apartments)
	return c.Status(fiber.StatusOK).JSON(resp)

}
