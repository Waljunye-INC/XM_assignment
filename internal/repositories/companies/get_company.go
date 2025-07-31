package companiesrepository

import (
	"context"

	"XM_assignment/internal/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/friendsofgo/errors"
)

func (r *repository) GetCompany(ctx context.Context, uuid string) (domain.Company, error) {
	qb := sq.Select(
		"uuid",
		"name",
		"description",
		"employees_count",
		"is_registered",
		"type").
		From("companies").
		Where(sq.Eq{"uuid": uuid})
	query, args, err := qb.ToSql()
	if err != nil {
		return domain.Company{}, errors.Wrap(err, "build query")
	}
	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return domain.Company{}, errors.Wrap(row.Err(), "query row")
	}

	var result domain.Company
	err = row.Scan(
		&result.UUID,
		&result.Name,
		&result.Description,
		&result.EmployeesCount,
		&result.IsRegistered,
		&result.Type,
	)
	if err != nil {
		return domain.Company{}, errors.Wrap(err, "row scan")
	}

	return result, nil
}
