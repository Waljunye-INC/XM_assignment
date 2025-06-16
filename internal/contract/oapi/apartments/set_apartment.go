package apartments

import (
	"OMS_assignment/internal/contract/oapi/apartments/dto"
	"OMS_assignment/internal/repository"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func (ac *Contract) SetApartment(c *fiber.Ctx) error {
	var request dto.Apartment

	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			errorField: err.Error(),
		})
	}

	apartment := dto.RequestApartmentToApartment(request)
	ctx := c.Context()

	err = ac.apartmentsUsecase.SetApartment(ctx, apartment)
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
