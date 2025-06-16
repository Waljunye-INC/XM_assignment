package apartments

import "context"

func (usc *apartmentsUsecase) DeleteApartment(ctx context.Context, id int64) error {
	err := usc.apartmentsRepository.DeleteApartment(ctx, id)
	if err != nil {
		// logging/business metrics if we need
		return err
	}

	return nil
}
