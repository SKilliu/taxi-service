package db

import (
	"github.com/SKilliu/taxi-service/db/models"
	"github.com/SKilliu/taxi-service/server/dto"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type OrdersQ interface {
	Insert(order models.Order) error
	Update(order models.Order) error
	Delete(order models.Order) error
	GetWithAvailableStatus() ([]models.Order, error)
	GetByID(oid string) (models.Order, error)
}

type OrdersWrapper struct {
	parent *DB
}

func (d *DB) OrdersQ() OrdersQ {
	return &OrdersWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (o *OrdersWrapper) Insert(order models.Order) error {
	return o.parent.db.Model(&order).Insert()
}

func (o *OrdersWrapper) Update(order models.Order) error {
	return o.parent.db.Model(&order).Update()
}

func (o *OrdersWrapper) Delete(order models.Order) error {
	return o.parent.db.Model(&order).Delete()
}

func (o *OrdersWrapper) GetWithAvailableStatus() ([]models.Order, error) {
	var res []models.Order
	err := o.parent.db.Select().From(models.OrdersTableName).Where(dbx.HashExp{"status": dto.StatusAvailable}).All(&res)
	return res, err
}

func (o *OrdersWrapper) GetByID(oid string) (models.Order, error) {
	var res models.Order
	err := o.parent.db.Select().From(models.OrdersTableName).Where(dbx.HashExp{"id": oid}).One(&res)
	return res, err
}
