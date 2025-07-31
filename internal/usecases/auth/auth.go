package auth

import (
	"XM_assignment/internal/domain"
	"context"
	"database/sql"
)

type repository interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	GetCredsByUsername(ctx context.Context, username string) (domain.Credentials, error)
	CreateCredsPair(ctx context.Context, credentials domain.Credentials, tx *sql.Tx) error
}

type useCase struct {
	repo   repository
	jwtKey string
}

func NewUseCase(jwtKey string, repo repository) *useCase {
	return &useCase{
		repo:   repo,
		jwtKey: jwtKey,
	}
}
