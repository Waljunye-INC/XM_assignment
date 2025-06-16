package buildings

import (
	"OMS_assignment/internal/repository"
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (bc *Contract) DeleteBuilding(c *fiber.Ctx) error {
	idRaw := c.Params(IDParam)

	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			errorField: err.Error(),
		})
	}
	ctx := c.Context()

	err = bc.buildingsUsecase.DeleteBuilding(ctx, id)
	if err != nil {
		status := fiber.StatusInternalServerError
		switch {
		case errors.Is(err, sql.ErrNoRows):
			status = fiber.StatusNotFound
		case errors.Is(err, repository.ErrNoRowsAffected):
			status = fiber.StatusUnprocessableEntity
		}

		return c.Status(status).JSON(fiber.Map{
			errorField: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
