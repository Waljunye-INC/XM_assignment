package companiescontract

import (
	"XM_assignment/internal/contract/oapi/contracterrors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

func (c *Contract) DeleteCompany(w http.ResponseWriter, r *http.Request) {
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

	err = c.useCase.DeleteCompany(ctx, uid.String())
	if err != nil {
		slog.Error("failed to delete company",
			"err", err,
			"company_uuid", uid)
		err = render.Render(w, r, contracterrors.ErrInternal(err))
		if err != nil {
			slog.Error("failed to render error", "err", err)
		}
		return
	}

	render.Status(r, http.StatusOK)
}
