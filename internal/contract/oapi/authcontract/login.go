package authcontract

import (
	"XM_assignment/internal/contract/oapi/authcontract/dto"
	"XM_assignment/internal/contract/oapi/contracterrors"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func (c *Contract) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var creds dto.CredentialsRequest
	if err := render.Bind(r, &creds); err != nil {
		slog.Error("no username & password provided",
			"err", err)
		err = render.Render(w, r, contracterrors.ErrBadRequest(err))
		if err != nil {
			slog.Error("failed to render error", "err", err)
		}
		return
	}

	at, err := c.usc.Login(ctx, dto.CredentialsFromRequest(creds))
	if err != nil {
		err = render.Render(w, r, contracterrors.ErrInternal(err))
		return
	}

	render.JSON(w, r, struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: at,
	})

}
