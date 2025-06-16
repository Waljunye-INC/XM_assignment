package buildings

import "database/sql"

type buildingsRepository struct {
	*sql.DB
}

func NewBuildingsRepository(db *sql.DB) *buildingsRepository {
	return &buildingsRepository{db}
}
