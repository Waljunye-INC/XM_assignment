package companiescontract

import (
	"XM_assignment/internal/contract/oapi/companiescontract/dto"
	"XM_assignment/internal/contract/oapi/contracterrors"
	"github.com/google/uuid"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (c *Contract) GetCompany(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uuidRaw := chi.URLParam(r, "uuid")
	uid, err := uuid.Parse(uuidRaw)
	if err != nil {
		slog.Error("failed to parse company uuid",
			"err", err,
			"raw_uuid", uuidRaw)
		err = render.Render(w, r, contracterrors.ErrBadRequest(err))
		if err != nil {
			slog.Error("failed to render error", "err", err)
		}
		return
	}

	company, err := c.useCase.GetCompany(ctx, uid.String())
	if err != nil {
		slog.Error("get company",
			"err", err,
			"uuid", uid)
		err = render.Render(w, r, contracterrors.ErrNotFound(err))
		if err != nil {
			slog.Error("failed to render error", "err", err)
		}
		return
	}

	render.JSON(w, r, dto.CompanyResponseFromDomain(company))
}
