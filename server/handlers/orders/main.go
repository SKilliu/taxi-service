package orders

import (
	"github.com/SKilliu/taxi-service/config"
	"github.com/SKilliu/taxi-service/db"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	auth         *config.Authentication
	log          *logrus.Entry
	usersDB      db.UsersQ
	carsDB       db.CarsQ
	driverCarsDB db.DriverCarsQ
	ordersDB     db.OrdersQ
}

func New(db db.QInterface, cfg config.Config) Handler {
	return Handler{
		auth:         cfg.Authentication(),
		log:          cfg.Log(),
		usersDB:      db.UsersQ(),
		carsDB:       db.CarsQ(),
		driverCarsDB: db.DriverCarsQ(),
		ordersDB:     db.OrdersQ(),
	}
}
