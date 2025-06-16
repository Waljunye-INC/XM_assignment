package buildings

import (
	"OMS_assignment/internal/domain"
	"OMS_assignment/internal/repository/models"
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (repo *buildingsRepository) GetBuildings(ctx context.Context) ([]domain.Building, error) {
	bs, err := models.Buildings(
		qm.Select("id", "name", "address"),
		qm.OrderBy("id DESC")).All(ctx, repo.DB)
	if err != nil {
		return nil, err
	}

	result := make([]domain.Building, 0, len(bs))
	for _, building := range bs {
		result = append(result, domain.Building{
			ID:      building.ID,
			Name:    building.Name,
			Address: building.Address,
		})
	}

	return result, nil
}
