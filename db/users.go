package db

import (
	"simple-service/db/models"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type UsersQ interface {
	Insert(user models.User) error
	Update(user models.User) error
	GetByEmail(email string) (models.User, error)
	GetAll() ([]models.User, error)
	GetByID(uid string) (models.User, error)
}

type UsersWrapper struct {
	parent *DB
}

func (d *DB) UsersQ() UsersQ {
	return &UsersWrapper{
		parent: &DB{d.db.Clone()},
	}
}

func (u *UsersWrapper) Insert(user models.User) error {
	return u.parent.db.Model(&user).Insert()
}

func (u *UsersWrapper) Update(user models.User) error {
	return u.parent.db.Model(&user).Update()
}

func (u *UsersWrapper) GetByEmail(email string) (models.User, error) {
	var res models.User
	err := u.parent.db.Select().From(models.UsersTableName).Where(dbx.HashExp{"email": email}).One(&res)
	return res, err
}

func (u *UsersWrapper) GetAll() ([]models.User, error) {
	var res []models.User
	err := u.parent.db.Select().From(models.UsersTableName).All(&res)
	return res, err
}

func (u *UsersWrapper) GetByID(uid string) (models.User, error) {
	var res models.User
	err := u.parent.db.Select().From(models.UsersTableName).Where(dbx.HashExp{"id": uid}).One(&res)
	return res, err
}
