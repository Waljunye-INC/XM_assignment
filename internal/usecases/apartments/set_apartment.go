package apartments

import (
	"OMS_assignment/internal/domain"
	"context"
)

func (usc *apartmentsUsecase) SetApartment(ctx context.Context, apartment domain.Apartment) error {
	err := usc.apartmentsRepository.SetApartment(ctx, apartment)
	if err != nil {
		// logging/business metrics if we need
		return err
	}

	return nil
}
