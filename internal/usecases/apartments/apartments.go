package apartments

import (
	"OMS_assignment/internal/domain"
	"context"
)

//go:generate mockgen -source=apartments.go -destination=mocks/mock_apartments_repository.go
type apartmentsRepository interface {
	GetApartments(ctx context.Context) (apartments []domain.Apartment, err error)
	GetApartmentByID(ctx context.Context, id int64) (apartment domain.Apartment, err error)
	GetApartmentsByBuildingID(ctx context.Context, buildingID int64) (apartments []domain.Apartment, err error)
	SetApartment(ctx context.Context, apartment domain.Apartment) (err error)
	DeleteApartment(ctx context.Context, id int64) (err error)
}

type apartmentsUsecase struct {
	apartmentsRepository apartmentsRepository
}

func NewApartmentsUsecase(apartmentsRepository apartmentsRepository) *apartmentsUsecase {
	return &apartmentsUsecase{apartmentsRepository: apartmentsRepository}
}
