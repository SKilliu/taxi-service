package db

import (
	"github.com/SKilliu/taxi-service/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type CarsQ interface {
	Insert(user models.Car) error
	Update(user models.Car) error
	Delete(car models.Car) error
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
