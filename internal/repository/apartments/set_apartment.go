package apartments

import (
	"OMS_assignment/internal/domain"
	"OMS_assignment/internal/repository/models"
	"context"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (repo *apartmentsRepository) SetApartment(ctx context.Context, apartment domain.Apartment) error {
	var apartmentToSet models.Apartment
	apartmentToSet.ID = apartment.ID
	apartmentToSet.BuildingID = apartment.BuildingID
	apartmentToSet.Floor = int(apartment.Floor)
	apartmentToSet.Number = null.String{String: apartment.Number, Valid: true}
	apartmentToSet.SQMeters = apartment.SQMeters

	err := apartmentToSet.Upsert(
		ctx,
		repo.DB,
		true,
		[]string{"id"},
		boil.Whitelist("building_id", "number", "floor", "sq_meters"),
		boil.Infer())
	if err != nil {
		return err
	}

	return nil

}
