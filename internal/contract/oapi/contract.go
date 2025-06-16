package oapi

import (
	"OMS_assignment/internal/contract/oapi/apartments"
	"OMS_assignment/internal/contract/oapi/buildings"
	"crypto/tls"

	"OMS_assignment/libs/http_server"
	"OMS_assignment/libs/listeners"

	"github.com/gofiber/fiber/v2"
)

func New(tls *tls.Certificate, bc *buildings.Contract, ac *apartments.Contract) listeners.PortListener {
	cont := &contract{
		bc: bc,
		ac: ac,
	}

	app := fiber.New()
	builds := app.Group("/buildings")
	{
		builds.Get("/", cont.bc.GetBuildings)
		builds.Get("/:id", cont.bc.GetBuildingByID)
		builds.Post("/", cont.bc.SetBuilding)
		builds.Delete("/:id", cont.bc.DeleteBuilding)
	}

	apparts := app.Group("/apartments")
	{
		apparts.Get("/", cont.ac.GetApartments)
		apparts.Get("/:id", cont.ac.GetApartmentByID)
		apparts.Get("/building/:buildingId", cont.ac.GetApartmentsByBuildingID)
		apparts.Post("/", cont.ac.SetApartment)
		apparts.Delete("/:id", cont.ac.DeleteApartment)
	}

	return http_server.New(app.Handler(), tls, "open_api")
}

type contract struct {
	bc *buildings.Contract
	ac *apartments.Contract
}
