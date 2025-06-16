package domain

type Apartment struct {
	ID         int64
	BuildingID int64
	Number     string
	Floor      int32
	SQMeters   float32
}

type Building struct {
	ID      int64
	Name    string
	Address string
}
