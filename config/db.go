package config

import (
	"fmt"

	"github.com/SKilliu/taxi-service/db"
	"github.com/caarlos0/env"
	validation "github.com/go-ozzo/ozzo-validation"
)

// DB is store all info about database connection
type DB struct {
	Name     string `env:"TAXI_SERVICE_API_DATABASE_NAME,required"`
	Host     string `env:"TAXI_SERVICE_API_DATABASE_HOST,required"`
	Port     int    `env:"TAXI_SERVICE_API_DATABASE_PORT,required"`
	User     string `env:"TAXI_SERVICE_API_DATABASE_USER,required"`
	Password string `env:"TAXI_SERVICE_API_DATABASE_PASSWORD,required"`
	SSL      string `env:"TAXI_SERVICE_API_DATABASE_SSL"`
}

// Info returning compiled statement. Statement is uri string which is provide access to the db
func (d DB) Info() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSL,
	)
}

// DB creates new local conn of the database
func (c *ConfigImpl) DB() db.QInterface {
	if c.db != nil {
		return c.db
	}

	log := c.Log()

	c.Lock()
	defer c.Unlock()

	database := &DB{}
	if err := env.Parse(database); err != nil {
		log.WithError(err).Error("failed to get db data from env")
		panic(err)
	}

	err := database.validate()
	if err != nil {
		log.WithError(err).Error("failed to validate db client")
		panic(err)
	}

	repo, err := db.New(database.Info())
	if err != nil {
		log.WithError(err).Error("failed to setup db")
		panic(err)
	}

	c.db = repo

	return c.db
}

func (d DB) validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.Port, validation.Required),
		validation.Field(&d.User, validation.Required),
		validation.Field(&d.Password, validation.Required),
	)
}
