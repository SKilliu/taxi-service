package db

import (
	"github.com/SKilliu/taxi-service/db/models"
)

type OrdersQ interface {
	Insert(order models.Order) error
	Update(order models.Order) error
	Delete(order models.Order) error
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
