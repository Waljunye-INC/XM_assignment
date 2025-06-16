package apartments

import (
	"OMS_assignment/internal/domain"
	"OMS_assignment/internal/repository/models"
	"context"
)

func (repo *apartmentsRepository) GetApartmentByID(ctx context.Context, id int64) (domain.Apartment, error) {
	apartment, err := models.FindApartment(ctx, repo.DB, id)
	if err != nil {
		return domain.Apartment{}, err
	}
	result := domain.Apartment{
		ID:         apartment.ID,
		BuildingID: apartment.BuildingID,
		Number:     apartment.Number.String,
		Floor:      int32(apartment.Floor),
		SQMeters:   apartment.SQMeters,
	}

	return result, nil
}
