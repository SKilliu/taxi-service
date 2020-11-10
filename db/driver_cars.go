package db

import (
	"github.com/SKilliu/taxi-service/db/models"
)

type DriverCarsQ interface {
	Insert(user models.Car) error
	Update(user models.Car) error
	Delete(car models.Car) error
}

type DriverCarsWrapper struct {
	parent *DB
}

func (d *DB) DriverCarsQ() DriverCarsQ {
	return &DriverCarsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (dc *DriverCarsWrapper) Insert(car models.Car) error {
	return dc.parent.db.Model(&car).Insert()
}

func (dc *DriverCarsWrapper) Update(car models.Car) error {
	return dc.parent.db.Model(&car).Update()
}

func (dc *DriverCarsWrapper) Delete(car models.Car) error {
	return dc.parent.db.Model(&car).Delete()
}
