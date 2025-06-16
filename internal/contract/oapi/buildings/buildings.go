package buildings

import (
	"OMS_assignment/internal/domain"
	"context"
)

type buildingsUsecase interface {
	GetBuildings(ctx context.Context) ([]domain.Building, error)
	GetBuildingByID(ctx context.Context, id int64) (domain.Building, error)
	SetBuilding(ctx context.Context, building domain.Building) error
	DeleteBuilding(ctx context.Context, id int64) error
}

type Contract struct {
	buildingsUsecase buildingsUsecase
}

func NewBuildingsContract(buildingsUsecase buildingsUsecase) *Contract {
	return &Contract{buildingsUsecase: buildingsUsecase}
}
