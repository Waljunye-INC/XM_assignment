package companiesrepository

import (
	"context"
	"database/sql"

	"XM_assignment/internal/repositories"

	sq "github.com/Masterminds/squirrel"
	"github.com/friendsofgo/errors"
)

func (r *repository) DeleteCompany(ctx context.Context, uuid string, tx *sql.Tx) error {
	qb := sq.Delete("companies").
		Where(sq.Eq{"uuid": uuid})

	query, args, err := qb.ToSql()
	if err != nil {
		return errors.Wrap(err, "building query")
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "query row")
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
