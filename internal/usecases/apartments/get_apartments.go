package apartments

import (
	"OMS_assignment/internal/domain"
	"context"
)

func (usc *apartmentsUsecase) GetApartments(ctx context.Context) ([]domain.Apartment, error) {
	result, err := usc.apartmentsRepository.GetApartments(ctx)
	if err != nil {
		// logging/business metrics if we need
		return []domain.Apartment{}, err
	}

	return result, nil
}
