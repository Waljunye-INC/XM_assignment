package companies

import (
	"XM_assignment/internal/domain"
	"context"
	"database/sql"
)

type companiesRepository interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	CreateCompany(ctx context.Context, company domain.Company, tx *sql.Tx) (createdCompany domain.Company, err error)
	UpdateCompany(ctx context.Context, company domain.Company, tx *sql.Tx) (err error)
	GetCompany(ctx context.Context, uuid string) (company domain.Company, err error)
	DeleteCompany(ctx context.Context, uuid string, tx *sql.Tx) (err error)
}

type eventsListener interface {
	ProduceEvent(ctx context.Context, message domain.Event) error
}

type useCase struct {
	repo     companiesRepository
	listener eventsListener
}

func NewUseCase(rep companiesRepository, listener eventsListener) *useCase {
	return &useCase{
		repo:     rep,
		listener: listener,
	}
}
