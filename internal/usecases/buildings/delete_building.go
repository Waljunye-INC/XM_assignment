package buildings

import "context"

func (usc *buildingsUsecase) DeleteBuilding(ctx context.Context, id int64) error {
	err := usc.buildingsRepository.DeleteBuilding(ctx, id)
	if err != nil {
		// logging/business metrics if we need

		return err
	}

	return nil
}
