package buildings

import (
	"OMS_assignment/internal/contract/oapi/buildings/dto"
	"database/sql"
	"github.com/friendsofgo/errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (bc *Contract) GetBuildingByID(c *fiber.Ctx) error {
	idRaw := c.Params(IDParam)

	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors.Wrap(err, "parse param \"id\" to int64"),
		})
	}
	ctx := c.Context()

	building, err := bc.buildingsUsecase.GetBuildingByID(ctx, id)
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
	resp := dto.BuildingToResponse(building)
	return c.Status(fiber.StatusOK).JSON(resp)
}
