package config

import (
	"github.com/caarlos0/env"
)

// Authentication is store all info about authentication
type Authentication struct {
	VerifyKey string `env:"TAXI_SERVICE_API_AUTHENTICATION_SECRET,required"`
	Algorithm string `env:"TAXI_SERVICE_AUTHENTICATION_ALGORITHM" envDefault:"HS256"`
}

// Authentication returns Authentication config
func (c *ConfigImpl) Authentication() *Authentication {
	if c.jwt != nil {
		return c.jwt
	}

	c.Lock()
	defer c.Unlock()

	jwt := &Authentication{}
	if err := env.Parse(jwt); err != nil {
		panic(err)
	}

	c.jwt = jwt

	return c.jwt
}
