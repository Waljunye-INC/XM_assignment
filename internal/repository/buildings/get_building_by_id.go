package buildings

import (
	"OMS_assignment/internal/domain"
	"OMS_assignment/internal/repository/models"
	"context"
)

func (repo *buildingsRepository) GetBuildingByID(ctx context.Context, id int64) (domain.Building, error) {
	queried, err := models.FindBuilding(ctx, repo.DB, id)
	if err != nil {
		return domain.Building{}, err
	}

	result := domain.Building{
		ID:      queried.ID,
		Name:    queried.Name,
		Address: queried.Address,
	}

	return result, nil
}
