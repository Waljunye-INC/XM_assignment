package companies

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"log/slog"

	"XM_assignment/internal/domain"

	"github.com/friendsofgo/errors"
)

func (usc *useCase) CreateCompany(ctx context.Context, company domain.Company) (domain.Company, error) {
	tx, err := usc.repo.BeginTx(ctx, &sql.TxOptions{})
	defer func() {
		err = tx.Rollback()
		if err != nil {
			slog.Error("error during rollback transaction",
				"err", err)
		}
	}()
	if err != nil {
		return domain.Company{}, errors.Wrap(err, "create transaction")
	}

	company.UUID = uuid.New().String()

	createdCompany, err := usc.repo.CreateCompany(ctx, company, tx)
	if err != nil {
		return domain.Company{}, errors.Wrap(err, "create company")
	}

	err = tx.Commit()
	if err != nil {
		return domain.Company{}, errors.Wrap(err, "transaction commit")
	}

	err = usc.listener.ProduceEvent(ctx, domain.Event{
		Key:       createdCompany.UUID,
		Message:   createdCompany,
		Operation: "created",
	})
	if err != nil {
		slog.Warn("unable to send event", "err", err)
	}

	return createdCompany, nil
}
