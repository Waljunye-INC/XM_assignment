package buildings

import (
	"OMS_assignment/internal/domain"
	"context"
)

func (usc *buildingsUsecase) GetBuildings(ctx context.Context) ([]domain.Building, error) {
	result, err := usc.buildingsRepository.GetBuildings(ctx)
	if err != nil {
		// logging/business metrics if we need
		return nil, err
	}

	return result, nil
}
