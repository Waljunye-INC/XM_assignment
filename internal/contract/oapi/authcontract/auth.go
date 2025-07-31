package authcontract

import (
	"XM_assignment/internal/domain"
	"context"
)

type useCase interface {
	Login(ctx context.Context, creds domain.Credentials) (at string, err error)
	Register(ctx context.Context, creds domain.Credentials) (err error)
}

type Contract struct {
	usc useCase
}

func NewContract(usc useCase) *Contract {
	return &Contract{
		usc: usc,
	}
}
