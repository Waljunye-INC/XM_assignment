package companies

import (
	"XM_assignment/internal/domain"
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"log/slog"
)

func (usc *useCase) UpdateCompany(ctx context.Context, company domain.Company) error {
	tx, err := usc.repo.BeginTx(ctx, &sql.TxOptions{})
	defer func() {
		err = tx.Rollback()
		if err != nil {
			slog.Error("error during rollback transaction",
				"err", err)
		}
	}()
	if err != nil {
		return errors.Wrap(err, "create transaction")
	}
	err = usc.repo.UpdateCompany(ctx, company, tx)
	if err != nil {
		return errors.Wrap(err, "update company")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "transaction commit")
	}

	err = usc.listener.ProduceEvent(ctx, domain.Event{
		Key:       company.UUID,
		Message:   company,
		Operation: "update",
	})
	if err != nil {
		slog.Warn("unable to send event", "err", err)
	}

	return nil
}
