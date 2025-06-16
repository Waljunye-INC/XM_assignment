package dto

import "OMS_assignment/internal/domain"

func ApartmentToResponse(apartment domain.Apartment) Apartment {
	return Apartment{
		ID:         apartment.ID,
		BuildingID: apartment.BuildingID,
		Number:     apartment.Number,
		Floor:      apartment.Floor,
		SQMeters:   apartment.SQMeters,
	}
}

func ApartmentsToResponse(apartments []domain.Apartment) Apartments {
	result := make([]Apartment, 0, len(apartments))

	for _, apartment := range apartments {
		result = append(result, ApartmentToResponse(apartment))
	}

	return result
}

func RequestApartmentToApartment(body Apartment) domain.Apartment {
	return domain.Apartment{
		ID:         body.ID,
		BuildingID: body.BuildingID,
		Number:     body.Number,
		Floor:      body.Floor,
		SQMeters:   body.SQMeters,
	}
}
