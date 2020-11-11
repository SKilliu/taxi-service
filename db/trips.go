package db

import (
	"github.com/SKilliu/taxi-service/db/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type TripsQ interface {
	Insert(order models.Trip) error
	Update(order models.Trip) error
	Delete(order models.Trip) error
	GetByID(tid string) (models.Trip, error)
}

type TripsWrapper struct {
	parent *DB
}

func (d *DB) TripsQ() TripsQ {
	return &TripsWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (o *TripsWrapper) Insert(order models.Trip) error {
	return o.parent.db.Model(&order).Insert()
}

func (o *TripsWrapper) Update(order models.Trip) error {
	return o.parent.db.Model(&order).Update()
}

func (o *TripsWrapper) Delete(order models.Trip) error {
	return o.parent.db.Model(&order).Delete()
}

func (o *TripsWrapper) GetByID(tid string) (models.Trip, error) {
	var res models.Trip
	err := o.parent.db.Select().From(models.TripsTableName).Where(dbx.HashExp{"id": tid}).One(&res)
	return res, err
}
