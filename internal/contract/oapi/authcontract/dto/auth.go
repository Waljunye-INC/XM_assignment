package dto

import (
	"XM_assignment/internal/domain"
	"net/http"
)

type CredentialsRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CredentialsFromRequest(cr CredentialsRequest) domain.Credentials {
	return domain.Credentials{
		Username: cr.Username,
		Password: cr.Password,
	}
}

func (cr *CredentialsRequest) Bind(_ *http.Request) error {
	return nil
}
