package buildings

import (
	"OMS_assignment/internal/contract/oapi/buildings/dto"
	"OMS_assignment/internal/repository"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func (bc *Contract) SetBuilding(c *fiber.Ctx) error {
	var request dto.Building

	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			errorField: err.Error(),
		})
	}

	building := dto.RequestApartmentToBuilding(request)
	ctx := c.Context()

	err = bc.buildingsUsecase.SetBuilding(ctx, building)
	if err != nil {
		status := fiber.StatusInternalServerError
		switch {
		case errors.Is(err, repository.ErrNoRowsAffected):
			status = fiber.StatusUnprocessableEntity
		}

		return c.Status(status).JSON(fiber.Map{
			errorField: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
