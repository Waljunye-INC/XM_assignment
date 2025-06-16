package apartments

import (
	"OMS_assignment/internal/domain"
	"context"
)

func (usc *apartmentsUsecase) GetApartmentByID(ctx context.Context, id int64) (domain.Apartment, error) {
	result, err := usc.apartmentsRepository.GetApartmentByID(ctx, id)
	if err != nil {
		// logging/business metrics if we need
		return domain.Apartment{}, err
	}

	return result, nil
}
