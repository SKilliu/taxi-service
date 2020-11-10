package models

import (
	"time"
)

const OrdersTableName = "orders"

type Order struct {
	ID                        string    `db:"id"`
	DriverID                  string    `db:"driver_id"`
	ClientID                  string    `db:"client_id"`
	CarID                     string    `db:"car_id"`
	StartingPointLongitude    float64   `db:"starting_point_longitude"`
	StartingPointLatitude     float64   `db:"starting_point_latitude"`
	DestinationPointLongitude float64   `db:"destination_point_longitude"`
	DestinationPointLatitude  float64   `db:"destination_point_latitude"`
	DestinationLocation       string    `db:"destination_location"`
	Price                     float64   `db:"price"`
	Status                    string    `db:"status"`
	CreatedAt                 time.Time `db:"created_at"`
}

func (o Order) TableName() string {
	return OrdersTableName
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}
