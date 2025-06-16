package apartments

import (
	"OMS_assignment/internal/domain"
	"OMS_assignment/internal/repository/models"
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (repo *apartmentsRepository) GetApartments(ctx context.Context) ([]domain.Apartment, error) {
	as, err := models.Apartments(qm.Select("id", "building_id", "number", "floor", "sq_meters")).All(ctx, repo.DB)
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
			SQMeters:   float32(apartment.SQMeters),
		})
	}

	return result, nil
}
