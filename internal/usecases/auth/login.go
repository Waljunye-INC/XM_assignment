package auth

import (
	"XM_assignment/internal/domain"
	"XM_assignment/utils"
	"context"
	"github.com/friendsofgo/errors"
	"golang.org/x/crypto/bcrypt"
)

func (u *useCase) Login(ctx context.Context, creds domain.Credentials) (string, error) {
	actualCreds, err := u.repo.GetCredsByUsername(ctx, creds.Username)
	if err != nil {
		return "", errors.Wrap(err, "get actual creds")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(actualCreds.Password), []byte(creds.Password)); err != nil {
		return "", errors.Wrap(ErrUnauthorized, "wrong user + password pair")
	}

	token, err := utils.GenerateJWT(creds.Username, u.jwtKey)
	if err != nil {
		return "", errors.Wrap(err, "generating token")
	}

	return token, nil
}
