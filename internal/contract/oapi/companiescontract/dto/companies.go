package dto

import (
	"XM_assignment/internal/domain"
	"errors"
	"net/http"
)

type CompanyResponse struct {
	UUID           string             `json:"uuid"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	EmployeesCount int32              `json:"employees_count"`
	IsRegistered   bool               `json:"is_registered"`
	Type           domain.CompanyType `json:"type"`
}

func (c CompanyResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func CompanyResponseFromDomain(comp domain.Company) CompanyResponse {
	return CompanyResponse{
		UUID:           comp.UUID,
		Name:           comp.Name,
		Description:    comp.Description,
		EmployeesCount: comp.EmployeesCount,
		IsRegistered:   comp.IsRegistered,
		Type:           comp.Type,
	}
}

type CompanyRequest struct {
	UUID           string             `json:"uuid"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	EmployeesCount int32              `json:"employees_count"`
	IsRegistered   bool               `json:"is_registered"`
	Type           domain.CompanyType `json:"type"`
}

func (cr *CompanyRequest) Bind(r *http.Request) error {
	if cr == nil {
		return errors.New("missing required fields")
	}

	return nil
}

func CompanyFromRequest(comp CompanyRequest) domain.Company {
	return domain.Company{
		UUID:           comp.UUID,
		Name:           comp.Name,
		Description:    comp.Description,
		EmployeesCount: comp.EmployeesCount,
		IsRegistered:   comp.IsRegistered,
		Type:           comp.Type,
	}
}
