package apartments

import (
	"OMS_assignment/internal/domain"
	"OMS_assignment/internal/repository/models"
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (repo *apartmentsRepository) GetApartmentsByBuildingID(ctx context.Context, buildingID int64) ([]domain.Apartment, error) {
	as, err := models.Apartments(
		qm.Select("id", "building_id", "number", "floor", "sq_meters"),
		qm.Where("building_id = ?", buildingID),
		qm.OrderBy("id DESC"),
	).All(ctx, repo.DB)
	if err != nil {
		return nil, err
	}

	result := make([]domain.Apartment, 0, len(as))
	for _, apartment := range as {
		result = append(result, domain.Apartment{
			ID:         apartment.ID,
			BuildingID: apartment.BuildingID,
			Number:     apartment.Number.String,
			Floor:      int32(apartment.Floor),
			SQMeters:   apartment.SQMeters,
		})
	}
	return result, nil

}
