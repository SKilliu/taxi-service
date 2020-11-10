package models

const DriverCarsTableName = "driver_cars"

type DriverCars struct {
	ID       string `db:"id"`
	DriverID string `db:"driver_id"`
	CarID    string `db:"car_id"`
}

func (c DriverCars) TableName() string {
	return DriverCarsTableName
}
