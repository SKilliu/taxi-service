package models

const TripsTableName = "trips"

type Trip struct {
	ID                        string  `db:"id"`
	StartingPointLocation     string  `db:"starting_point_location"`
	StartingPointLongitude    float64 `db:"starting_point_longitude"`
	StartingPointLatitude     float64 `db:"starting_point_latitude"`
	DestinationPointLocation  string  `db:"destination_location"`
	DestinationPointLongitude float64 `db:"destination_point_longitude"`
	DestinationPointLatitude  float64 `db:"destination_point_latitude"`
	Distance                  float64 `db:"distance"`
}

func (t Trip) TableName() string {
	return TripsTableName
}
