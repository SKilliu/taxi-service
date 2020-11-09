package db

import dbx "github.com/go-ozzo/ozzo-dbx"

type QInterface interface {
	DBX() *dbx.DB
	UsersQ() UsersQ
}

// DB wraps dbx interface.
type DB struct {
	db *dbx.DB
}

// DBX returns db client.
func (d DB) DBX() *dbx.DB {
	return d.db
}

// New connection opening.
func New(link string) (QInterface, error) {
	db, err := dbx.Open("postgres", link)
	return &DB{db: db}, err
}
