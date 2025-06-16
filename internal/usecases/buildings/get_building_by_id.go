package buildings

import (
	"OMS_assignment/internal/domain"
	"context"
)

func (usc *buildingsUsecase) GetBuildingByID(ctx context.Context, id int64) (domain.Building, error) {
	result, err := usc.buildingsRepository.GetBuildingByID(ctx, id)
	if err != nil {
		// logging/business metrics if we need
		return domain.Building{}, err
	}

	return result, nil
}
