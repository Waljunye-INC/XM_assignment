package dto

type Apartment struct {
	ID         int64   `json:"id"`
	BuildingID int64   `json:"building_id"`
	Number     string  `json:"number" omitempty:"true"`
	Floor      int32   `json:"floor"`
	SQMeters   float32 `json:"sq_meters"`
}

type Apartments []Apartment
