package buildings

import (
	"OMS_assignment/internal/domain"
	"context"
)

//go:generate mockgen -source=buildings.go -destination=mocks/mock_buildings_repository.go
type buildingsRepository interface {
	GetBuildings(ctx context.Context) (buildings []domain.Building, err error)
	GetBuildingByID(ctx context.Context, id int64) (building domain.Building, err error)
	SetBuilding(ctx context.Context, building domain.Building) (err error)
	DeleteBuilding(ctx context.Context, id int64) (err error)
}

type buildingsUsecase struct {
	buildingsRepository buildingsRepository
}

func NewBuildingsUsecase(buildingsRepository buildingsRepository) *buildingsUsecase {
	return &buildingsUsecase{buildingsRepository: buildingsRepository}
}
