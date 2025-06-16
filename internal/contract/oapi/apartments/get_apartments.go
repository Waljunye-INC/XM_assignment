package apartments

import (
	"OMS_assignment/internal/contract/oapi/apartments/dto"
	"github.com/gofiber/fiber/v2"
)

func (ac *Contract) GetApartments(c *fiber.Ctx) error {
	ctx := c.Context()

	apartments, err := ac.apartmentsUsecase.GetApartments(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			errorField: err.Error(),
		})
	}

	resp := dto.ApartmentsToResponse(apartments)
	return c.Status(fiber.StatusOK).JSON(resp)
}
