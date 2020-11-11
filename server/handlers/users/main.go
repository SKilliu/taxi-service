package users

import (
	"github.com/SKilliu/taxi-service/config"
	"github.com/SKilliu/taxi-service/db"
	"github.com/SKilliu/taxi-service/s3"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	auth    *config.Authentication
	log     *logrus.Entry
	usersDB db.UsersQ
	s3      *s3.Client
}

func New(db db.QInterface, cfg config.Config) Handler {
	return Handler{
		auth:    cfg.Authentication(),
		log:     cfg.Log(),
		usersDB: db.UsersQ(),
		s3:      cfg.S3(),
	}
}
