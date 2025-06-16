package buildings

import (
	"OMS_assignment/internal/repository"
	"context"

	"OMS_assignment/internal/repository/models"
)

func (repo *buildingsRepository) DeleteBuilding(ctx context.Context, id int64) error {
	buildingToDelete, err := models.FindBuilding(ctx, repo.DB, id)
	if err != nil {
		return err
	}

	rowsAff, err := buildingToDelete.Delete(ctx, repo.DB)
	if err != nil {
		return err
	}

	if rowsAff == 0 {
		return repository.ErrNoRowsAffected
	}

	return nil
}
