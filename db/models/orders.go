package models

import (
	"time"
)

const OrdersTableName = "orders"

type Order struct {
	ID             string    `db:"id"`
	DriverID       string    `db:"driver_id"`
	ClientID       string    `db:"client_id"`
	CarID          string    `db:"car_id"`
	TripID         string    `db:"trip_id"`
	Price          float64   `db:"price"`
	Status         string    `db:"status"`
	CarArrivalTime time.Time `db:"car_arrival_time"`
	CreatedAt      time.Time `db:"created_at"`
}

func (o Order) TableName() string {
	return OrdersTableName
}
