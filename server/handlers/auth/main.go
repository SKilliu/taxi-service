package auth

import (
	"simple-service/config"
	"simple-service/db"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	auth    *config.Authentication
	log     *logrus.Entry
	usersDB db.UsersQ
}

func New(db db.QInterface, cfg config.Config) Handler {
	return Handler{
		auth:    cfg.Authentication(),
		log:     cfg.Log(),
		usersDB: db.UsersQ(),
	}
}
