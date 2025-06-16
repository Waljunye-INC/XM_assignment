package buildings

import (
	"OMS_assignment/internal/domain"
	"OMS_assignment/internal/repository/models"
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (repo *buildingsRepository) SetBuilding(ctx context.Context, building domain.Building) error {
	var buildingToSet models.Building
	buildingToSet.ID = building.ID
	buildingToSet.Name = building.Name
	buildingToSet.Address = building.Address

	err := buildingToSet.Upsert(
		ctx,
		repo.DB,
		true,
		[]string{"id"},
		boil.Whitelist("name", "address"),
		boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
