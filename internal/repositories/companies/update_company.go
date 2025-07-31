package companiesrepository

import (
	"context"
	"database/sql"

	"XM_assignment/internal/domain"
	"XM_assignment/internal/repositories"

	sq "github.com/Masterminds/squirrel"
	"github.com/friendsofgo/errors"
)

func (r *repository) UpdateCompany(ctx context.Context, company domain.Company, tx *sql.Tx) (err error) {
	qb := sq.
		Update("companies").
		Set("name", company.Name).
		Set("description", company.Description).
		Set("employees_count", company.EmployeesCount).
		Set("is_registered", company.IsRegistered).
		Set("type", company.Type).
		Where(sq.Eq{"uuid": company.UUID})
	query, args, err := qb.ToSql()
	if err != nil {
		return errors.Wrap(err, "build query")
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "exec query")
	}
	var ra int64
	ra, err = result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "scanning result")
	}

	if ra == 0 {
		return repositories.ErrNoRowsAffected
	}

	return nil
}
