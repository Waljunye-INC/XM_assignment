package oapi

import (
	"crypto/tls"
	"github.com/go-chi/chi/v5"

	"XM_assignment/internal/contract/oapi/authcontract"
	"XM_assignment/internal/contract/oapi/companiescontract"

	"XM_assignment/libs/http_server"
	"XM_assignment/libs/listeners"
)

func New(tls *tls.Certificate, jwtKey string, cc *companiescontract.Contract, ac *authcontract.Contract) listeners.PortListener {
	cont := &contract{
		companyContract: cc,
		authContract:    ac,
	}

	app := chi.NewRouter()
	app.Route("/companies", func(compRoute chi.Router) {
		compRoute.Get("/{uuid}", cont.companyContract.GetCompany)
		compRoute.Group(func(protecredCompRoute chi.Router) {
			protecredCompRoute.Use(JWTAuthMiddleware(jwtKey))
			protecredCompRoute.Post("/", cont.companyContract.CreateCompany)
			protecredCompRoute.Put("/", cont.companyContract.UpdateCompany)
			protecredCompRoute.Delete("/{uuid}", cont.companyContract.DeleteCompany)
		})
	})

	app.Route("/auth", func(authRoute chi.Router) {
		authRoute.Post("/register", cont.authContract.Register)
		authRoute.Put("/login", cont.authContract.Login)
	})

	return http_server.New(app, tls, "open_api")
}

type contract struct {
	companyContract *companiescontract.Contract
	authContract    *authcontract.Contract
}
