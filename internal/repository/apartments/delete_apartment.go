package apartments

import (
	"OMS_assignment/internal/repository"
	"OMS_assignment/internal/repository/models"
	"context"
)

func (repo *apartmentsRepository) DeleteApartment(ctx context.Context, id int64) error {
	apartment, err := models.FindApartment(ctx, repo.DB, id)
	if err != nil {
		return err
	}
	rowsAff, err := apartment.Delete(ctx, repo.DB)
	if err != nil {
		return err
	}
	if rowsAff == 0 {
		return repository.ErrNoRowsAffected
	}

	return nil
}
