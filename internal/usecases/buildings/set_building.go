package buildings

import (
	"OMS_assignment/internal/domain"
	"context"
)

func (usc *buildingsUsecase) SetBuilding(ctx context.Context, building domain.Building) error {
	err := usc.buildingsRepository.SetBuilding(ctx, building)
	if err != nil {
		// logging/business metrics if we need

		return err
	}

	return nil
}
