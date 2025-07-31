package companiescontract

import (
	"context"

	"XM_assignment/internal/domain"
)

//go:generate mockgen -source=companies.go -destination=mocks/companies.go
type companiesUseCase interface {
	GetCompany(ctx context.Context, uuid string) (company domain.Company, err error)
	UpdateCompany(ctx context.Context, updatedCompany domain.Company) (err error)
	CreateCompany(ctx context.Context, company domain.Company) (createdCompany domain.Company, err error)
	DeleteCompany(ctx context.Context, uuid string) (err error)
}

type Contract struct {
	useCase companiesUseCase
}

func NewContract(usc companiesUseCase) *Contract {
	return &Contract{
		useCase: usc,
	}
}
