package authrepository

import (
	"XM_assignment/internal/domain"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/friendsofgo/errors"
)

func (r *repository) CreateCredsPair(ctx context.Context, credentials domain.Credentials, tx *sql.Tx) error {
	qb := sq.Insert("credentials").
		Columns("username", "password").
		Values(credentials.Username, credentials.Password)

	query, args, err := qb.ToSql()
	if err != nil {
		return errors.Wrap(err, "building query")
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "exec query")
	}
	return err
}
