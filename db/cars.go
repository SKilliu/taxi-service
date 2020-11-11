package db

import (
	"fmt"

	"github.com/SKilliu/taxi-service/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

const getAvailableCar = "SELECT * FROM cars WHERE status = 'available' AND id = " +
	"(SELECT car_id FROM driver_cars WHERE driver_id = '%s' LIMIT 1) LIMIT 1);"

type CarsQ interface {
	Insert(user models.Car) error
	Update(user models.Car) error
	Delete(car models.Car) error
	GetByNumber(number string) (models.Car, error)
	GetAvailableCar(driverID string) (models.Car, error)
	GetByID(cid string) (models.Car, error)
}

type CarsWrapper struct {
	parent *DB
}

func (d *DB) CarsQ() CarsQ {
	return &CarsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (c *CarsWrapper) Insert(car models.Car) error {
	return c.parent.db.Model(&car).Insert()
}

func (c *CarsWrapper) Update(car models.Car) error {
	return c.parent.db.Model(&car).Update()
}

func (c *CarsWrapper) Delete(car models.Car) error {
	return c.parent.db.Model(&car).Delete()
}

func (c *CarsWrapper) GetByID(cid string) (models.Car, error) {
	var res models.Car
	err := c.parent.db.Select().From(models.CarsTableName).Where(dbx.HashExp{"id": cid}).One(&res)
	return res, err
}

func (c *CarsWrapper) GetByNumber(number string) (models.Car, error) {
	var res models.Car
	err := c.parent.db.Select().From(models.CarsTableName).Where(dbx.HashExp{"number": number}).One(&res)
	return res, err
}

func (c *CarsWrapper) GetAvailableCar(driverID string) (models.Car, error) {
	var res models.Car
	err := c.parent.db.NewQuery(fmt.Sprintf(getAvailableCar, driverID)).One(&res)
	return res, err
}
