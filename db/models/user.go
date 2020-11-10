package models

const UsersTableName = "users"

type User struct {
	ID              string `db:"id"`
	Name            string `db:"name"`
	AccountType     string `db:"account_type"`
	HashedPassword  string `db:"hashed_password"`
	Email           string `db:"email"`
	ProfileImageUrl string `db:"profile_image_url"`
}

func (u User) TableName() string {
	return UsersTableName
}
