package auth

import (
	"XM_assignment/internal/domain"
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func (usc *useCase) Register(ctx context.Context, creds domain.Credentials) error {
	tx, err := usc.repo.BeginTx(ctx, &sql.TxOptions{})
	defer func() {
		err = tx.Rollback()
		if err != nil {
			slog.Error("error during rollback transaction",
				"err", err)
		}
	}()

	if creds.Username == "" || creds.Password == "" {
		return errors.New("empty username or password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "hashing password")
	}
	err = usc.repo.CreateCredsPair(ctx, domain.Credentials{
		Username: creds.Username,
		Password: string(hashedPassword),
	}, tx)

	if err != nil {
		return errors.Wrap(err, "creating pair")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "transaction commit")
	}

	return nil
}
