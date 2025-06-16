package apartments

import (
	"OMS_assignment/internal/contract/oapi/apartments/dto"
	"database/sql"
	"github.com/friendsofgo/errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (ac *Contract) GetApartmentByID(c *fiber.Ctx) error {
	idRaw := c.Params(IDParam)

	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			errorField: err.Error(),
		})
	}
	ctx := c.Context()

	apartment, err := ac.apartmentsUsecase.GetApartmentByID(ctx, id)
	if err != nil {
		status := fiber.StatusInternalServerError
		switch {
		case errors.Is(err, sql.ErrNoRows):
			status = fiber.StatusNotFound
		}
		return c.Status(status).JSON(fiber.Map{
			errorField: err.Error(),
		})
	}

	resp := dto.ApartmentToResponse(apartment)
	return c.Status(fiber.StatusOK).JSON(resp)
}
