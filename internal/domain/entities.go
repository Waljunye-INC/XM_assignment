package domain

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type Company struct {
	UUID           string
	Name           string
	Description    string
	EmployeesCount int32
	IsRegistered   bool
	Type           CompanyType
}

const (
	CompanyTypeCorporation        = "corporation"
	CompanyTypeNonProfit          = "non_profit"
	CompanyTypeCooperative        = "cooperative"
	CompanyTypeSoleProprietorship = "sole_proprietorship"
)

type CompanyType string

var companyTypesMap = map[string]struct{}{
	CompanyTypeCorporation:        {},
	CompanyTypeNonProfit:          {},
	CompanyTypeCooperative:        {},
	CompanyTypeSoleProprietorship: {},
}

func (ct CompanyType) IsValid() bool {
	_, exists := companyTypesMap[string(ct)]
	return exists
}

func CompanyTypeFromString(str string) (CompanyType, error) {
	if _, valid := companyTypesMap[str]; valid {
		return CompanyType(str), nil
	}
	return "", errors.New("wrong company type")
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Credentials struct {
	Username string
	Password string
}
