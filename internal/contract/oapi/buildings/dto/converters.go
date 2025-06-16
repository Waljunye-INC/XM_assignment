package dto

import "OMS_assignment/internal/domain"

func BuildingToResponse(building domain.Building) Building {
	return Building{
		ID:      building.ID,
		Name:    building.Name,
		Address: building.Address,
	}
}

func BuildingsToResponse(buildings []domain.Building) Buildings {
	result := make([]Building, 0, len(buildings))

	for _, building := range buildings {
		result = append(result, BuildingToResponse(building))
	}

	return result
}

func RequestApartmentToBuilding(body Building) domain.Building {
	return domain.Building{
		ID:      body.ID,
		Name:    body.Name,
		Address: body.Address,
	}
}
