package companiescontract

import (
	"XM_assignment/internal/contract/oapi/contracterrors"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"

	"XM_assignment/internal/contract/oapi/companiescontract/dto"
)

func (c *Contract) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	companyRequest := dto.CompanyRequest{}
	if err := render.Bind(r, &companyRequest); err != nil {
		slog.Error("no company provided",
			"err", err)
		err = render.Render(w, r, contracterrors.ErrBadRequest(err))
		if err != nil {
			slog.Error("failed to render error", "err", err)
		}
		return
	}

	err := c.useCase.UpdateCompany(ctx, dto.CompanyFromRequest(companyRequest))
	if err != nil {
		compBytes, _ := json.Marshal(companyRequest)
		slog.Error("failed to update company",
			"err", err,
			"request_body", string(compBytes))
		err = render.Render(w, r, contracterrors.ErrInternal(err))
		if err != nil {
			slog.Error("failed to render error", "err", err)
		}
		return
	}

	render.Status(r, http.StatusOK)
}
