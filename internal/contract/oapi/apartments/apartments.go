package apartments

import (
	"OMS_assignment/internal/domain"
	"context"
)

type apartmentsUsecase interface {
	GetApartments(ctx context.Context) (apartments []domain.Apartment, err error)
	GetApartmentByID(ctx context.Context, id int64) (apartment domain.Apartment, err error)
	GetApartmentsByBuildingID(ctx context.Context, buildingID int64) (apartments []domain.Apartment, err error)
	SetApartment(ctx context.Context, apartment domain.Apartment) (err error)
	DeleteApartment(ctx context.Context, id int64) (err error)
}

type Contract struct {
	apartmentsUsecase apartmentsUsecase
}

func NewApartmentsContract(apartmentsUsecase apartmentsUsecase) *Contract {
	return &Contract{apartmentsUsecase: apartmentsUsecase}
}
