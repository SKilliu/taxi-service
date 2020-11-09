package server

import (
	"github.com/SKilliu/taxi-service/config"
	"github.com/SKilliu/taxi-service/db"
	"github.com/SKilliu/taxi-service/server/handlers/auth"
	"github.com/SKilliu/taxi-service/server/handlers/users"
)

// Provider persist handlers service structures.
type Provider struct {
	UsersHandler users.Handler
	AuthHandler  auth.Handler
}

// NewProvider is provider constructor.
func NewProvider(cfg config.Config, db db.QInterface) *Provider {
	return &Provider{
		UsersHandler: users.New(db, cfg),
		AuthHandler:  auth.New(db, cfg),
	}
}
