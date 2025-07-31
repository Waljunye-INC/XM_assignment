package companiesrepository

import (
	"context"
	"database/sql"

	"XM_assignment/internal/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/friendsofgo/errors"
)

func (r *repository) CreateCompany(ctx context.Context, company domain.Company, tx *sql.Tx) (domain.Company, error) {
	qb := sq.Insert("companies").
		Columns("uuid", "name", "description", "employees_count", "is_registered", "type").
		Values(company.UUID, company.Name, company.Description, company.EmployeesCount, company.IsRegistered, company.Type)
	query, args, err := qb.ToSql()
	if err != nil {
		return domain.Company{}, errors.Wrap(err, "building query")
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return domain.Company{}, errors.Wrap(err, "exec")
	}

	var insertedCompany domain.Company
	err = tx.QueryRowContext(ctx, "SELECT uuid, name, description, employees_count, is_registered, type FROM companies WHERE uuid = ?", company.UUID).
		Scan(&insertedCompany.UUID, &insertedCompany.Name, &insertedCompany.Description, &insertedCompany.EmployeesCount, &insertedCompany.IsRegistered, &insertedCompany.Type)
	if err != nil {
		return domain.Company{}, errors.Wrap(err, "scanning inserted")
	}

	return insertedCompany, nil
}
