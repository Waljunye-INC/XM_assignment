package apartments

import "database/sql"

type apartmentsRepository struct {
	*sql.DB
}

func NewApartmentsRepository(db *sql.DB) *apartmentsRepository {
	return &apartmentsRepository{db}
}
