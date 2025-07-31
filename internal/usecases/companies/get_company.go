package companies

import (
	"XM_assignment/internal/domain"
	"context"
	"github.com/friendsofgo/errors"
)

func (usc *useCase) GetCompany(ctx context.Context, uuid string) (domain.Company, error) {
	company, err := usc.repo.GetCompany(ctx, uuid)
	if err != nil {
		return domain.Company{}, errors.Wrap(err, "get company")
	}
	return company, nil
}
