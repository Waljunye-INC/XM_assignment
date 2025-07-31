package companies

import (
	"XM_assignment/internal/domain"
	"context"
	"database/sql"
	"github.com/friendsofgo/errors"
	"log/slog"
)

func (usc *useCase) DeleteCompany(ctx context.Context, uuid string) error {
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
	err = usc.repo.DeleteCompany(ctx, uuid, tx)
	if err != nil {
		return errors.Wrap(err, "delete company")
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "transaction commit")
	}
	err = usc.listener.ProduceEvent(ctx, domain.Event{
		Key:       uuid,
		Operation: "deleted",
	})
	if err != nil {
		slog.Warn("unable to send event", "err", err)
	}

	return nil
}
