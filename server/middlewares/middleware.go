package middlewares

import "simple-service/config"

type Middleware struct {
	auth *config.Authentication
}

func New(cfg config.Config) *Middleware {
	return &Middleware{
		auth: cfg.Authentication(),
	}
}
