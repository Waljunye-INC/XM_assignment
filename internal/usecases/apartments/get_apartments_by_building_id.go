package apartments

import (
	"OMS_assignment/internal/domain"
	"context"
)

func (usc *apartmentsUsecase) GetApartmentsByBuildingID(ctx context.Context, buildingID int64) ([]domain.Apartment, error) {
	result, err := usc.apartmentsRepository.GetApartmentsByBuildingID(ctx, buildingID)
	if err != nil {
		// logging/business metrics if we need
		return []domain.Apartment{}, err
	}

	return result, nil
}
