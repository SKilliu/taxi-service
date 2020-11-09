package middlewares

import "github.com/SKilliu/taxi-service/config"

type Middleware struct {
	auth *config.Authentication
}

func New(cfg config.Config) *Middleware {
	return &Middleware{
		auth: cfg.Authentication(),
	}
}
