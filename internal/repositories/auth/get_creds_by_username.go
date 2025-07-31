package authrepository

import (
	"XM_assignment/internal/domain"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/friendsofgo/errors"
)

func (r *repository) GetCredsByUsername(ctx context.Context, username string) (domain.Credentials, error) {
	qb := sq.Select("username", "password").
		From("credentials").
		Where(sq.Eq{"username": username})

	query, args, err := qb.ToSql()
	if err != nil {
		return domain.Credentials{}, errors.Wrap(err, "building query")
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	if row.Err() != nil {
		return domain.Credentials{}, errors.Wrap(row.Err(), "query row")
	}
	var result domain.Credentials
	err = row.Scan(
		&result.Username,
		&result.Password,
	)

	if err != nil {
		return domain.Credentials{}, errors.Wrap(err, "scan result")
	}

	return result, nil
}
